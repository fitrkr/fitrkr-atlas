package commands

import (
	"context"
	"fmt"
	"slices"

	"github.com/cheezecakee/logr"

	"github.com/fitrkr/atlas/internal/core/domain/exercise"
)

type ExerciseRelation interface {
	Add(ctx context.Context, exerciseID int) error
	Remove(ctx context.Context) error
}

type ExerciseBuilder struct {
	ctx         context.Context
	exerciseID  *int
	equipmentID *int
	relations   []ExerciseRelation
}

func NewExerciseBuilder(ctx context.Context, exerciseID int) *ExerciseBuilder {
	return &ExerciseBuilder{
		ctx:        ctx,
		exerciseID: &exerciseID,
		relations:  []ExerciseRelation{},
	}
}

func (b *ExerciseBuilder) WithAlias(add []exercise.Alias, remove []int) *ExerciseBuilder {
	b.relations = append(b.relations, &AliasRelation{
		ToAdd:    add,
		ToRemove: remove,
	})
	return b
}

func (b *ExerciseBuilder) WithMuscles(add []MuscleReq, remove []int) *ExerciseBuilder {
	b.relations = append(b.relations, &MuscleRelation{
		ToAdd:    add,
		ToRemove: remove,
	})
	return b
}

func (b *ExerciseBuilder) WithCategories(add []int, remove []int) *ExerciseBuilder {
	b.relations = append(b.relations, &CategoryRelation{
		ToAdd:    add,
		ToRemove: remove,
	})
	return b
}

func (b *ExerciseBuilder) WithAttachments(add []int, remove []int) *ExerciseBuilder {
	b.relations = append(b.relations, &AttachmentRelation{
		ToAdd:       add,
		ToRemove:    remove,
		equipmentID: b.equipmentID,
	})
	return b
}

func (b *ExerciseBuilder) Execute() error {
	for _, relation := range b.relations {
		// Remove first (safe to call if empty)
		if err := relation.Remove(b.ctx); err != nil {
			return fmt.Errorf("failed to remove relation: %w", err)
		}

		if b.exerciseID != nil {
			if err := relation.Add(b.ctx, *b.exerciseID); err != nil {
				return fmt.Errorf("failed to add relation: %w", err)
			}
		}
	}

	return nil
}

// ---------- Relation Implementations ----------

type AliasRelation struct {
	ToAdd    []exercise.Alias
	ToRemove []int
}

func (r *AliasRelation) Add(ctx context.Context, exerciseID int) error {
	if len(r.ToAdd) == 0 {
		return nil
	}

	var aliases []exercise.Alias
	for _, i := range r.ToAdd {
		// Check if an alias exists with the same name
		if _, err := read.Exercise.Alias.GetByName(ctx, i.Name); err == nil {
			return fmt.Errorf("alias already exists")
		}
		alias, err := exercise.NewAlias(exerciseID, i.Name, i.LanguageCode)
		if err != nil {
			return fmt.Errorf("failed to create exercise alias: %w", err)
		}
		aliases = append(aliases, alias)
	}

	return write.Exercise.Alias.Add(ctx, aliases)
}

func (r *AliasRelation) Remove(ctx context.Context) error {
	if len(r.ToRemove) == 0 {
		return nil
	}
	return write.Exercise.Alias.Delete(ctx, r.ToRemove)
}

type MuscleReq struct {
	MuscleID   int    `json:"muscle_id"`
	Activation string `json:"activation"`
}

type MuscleRelation struct {
	ToAdd    []MuscleReq
	ToRemove []int
}

func (r *MuscleRelation) Add(ctx context.Context, exerciseID int) error {
	if len(r.ToAdd) == 0 {
		return nil
	}

	var muscles []exercise.ExerciseMuscle
	for _, req := range r.ToAdd {
		level, err := exercise.NewActivationLevel(req.Activation)
		if err != nil {
			return fmt.Errorf("failed to create activation level: %w", err)
		}

		muscle, err := exercise.NewExerciseMuscle(exerciseID, req.MuscleID, level.ToString())
		if err != nil {
			return fmt.Errorf("failed to create exercise muscle: %w", err)
		}
		muscles = append(muscles, muscle)
	}

	return write.Exercise.Muscle.Add(ctx, muscles)
}

func (r *MuscleRelation) Remove(ctx context.Context) error {
	if len(r.ToRemove) == 0 {
		return nil
	}
	return write.Exercise.Muscle.Delete(ctx, r.ToRemove)
}

type CategoryRelation struct {
	ToAdd    []int
	ToRemove []int
}

func (r *CategoryRelation) Add(ctx context.Context, exerciseID int) error {
	if len(r.ToAdd) == 0 {
		return nil
	}

	var categories []exercise.ExerciseCategory
	for _, categoryID := range r.ToAdd {
		category, err := exercise.NewExerciseCategory(exerciseID, categoryID)
		if err != nil {
			return fmt.Errorf("failed to create exercise category: %w", err)
		}
		categories = append(categories, category)
	}

	return write.Exercise.Category.Add(ctx, categories)
}

func (r *CategoryRelation) Remove(ctx context.Context) error {
	if len(r.ToRemove) == 0 {
		return nil
	}
	return write.Exercise.Category.Delete(ctx, r.ToRemove)
}

type AttachmentRelation struct {
	ToAdd       []int
	ToRemove    []int
	equipmentID *int
}

func (r *AttachmentRelation) Add(ctx context.Context, exerciseID int) error {
	if len(r.ToAdd) == 0 || r.equipmentID == nil {
		return nil
	}

	var eqAttachments []int
	// Check if attachment selected is compatible with equipment
	equipmentAttachments, err := read.Equipment.Attachment.GetByEquipmentID(ctx, *r.equipmentID)
	if err != nil {
		return fmt.Errorf("failed to read equipment attachment: %w", err)
	}

	for _, i := range equipmentAttachments {
		eqAttachments = append(eqAttachments, i.AttachmentID)
	}

	var attachments []exercise.ExerciseAttachment
	for _, attachmentID := range r.ToAdd {
		if !slices.Contains(eqAttachments, attachmentID) {
			logr.Get().Infof("attachment not compatible with equipment %d:%d", attachmentID, *r.equipmentID)
			continue
		}

		attachment, err := exercise.NewExerciseAttachment(exerciseID, &attachmentID)
		if err != nil {
			return fmt.Errorf("failed to create exercise attachment: %w", err)
		}
		attachments = append(attachments, *attachment)
	}

	return write.Exercise.Attachment.Add(ctx, attachments)
}

func (r *AttachmentRelation) Remove(ctx context.Context) error {
	if len(r.ToRemove) == 0 {
		return nil
	}
	return write.Exercise.Attachment.Delete(ctx, r.ToRemove)
}
