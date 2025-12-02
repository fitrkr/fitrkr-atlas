package commands

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/view"
)

type BuildViewCommand struct {
	ExerciseID int  `json:"exercise_id"`
	Create     bool `json:"create"`
}

type BuildViewResp struct{}

func (cmd *BuildViewCommand) Handle(ctx context.Context) (any, error) {
	ex, err := read.Exercise.GetByID(ctx, cmd.ExerciseID)
	if err != nil {
		return nil, err
	}

	equipment, err := cmd.BuildEquipment(ctx, ex.EquipmentID)
	if err != nil {
		return nil, err
	}

	muscles, err := cmd.BuildMuscle(ctx)
	if err != nil {
		return nil, err
	}

	categories, err := cmd.BuildCategory(ctx)
	if err != nil {
		return nil, err
	}

	aliases, err := cmd.BuildAlias(ctx)
	if err != nil {
		return nil, err
	}

	v := NewViewBuilder(ex).
		WithEquipment(equipment).
		WithMuscles(muscles).
		WithCategory(categories).
		WithAliases(aliases).
		Build()

	if !cmd.Create {
		err = write.View.Update(ctx, *v)
		if err != nil {
			return BuildViewResp{}, fmt.Errorf("failed update exercise view: %w", err)
		}

		return BuildViewResp{}, nil
	}
	_, err = write.View.Add(ctx, *v)
	if err != nil {
		return BuildViewResp{}, fmt.Errorf("failed insert exercise view: %w", err)
	}

	return BuildViewResp{}, nil
}

type ViewBuilder struct {
	view *view.View
}

func NewViewBuilder(ex *exercise.Exercise) *ViewBuilder {
	return &ViewBuilder{
		view: &view.View{
			ID:          *ex.ID,
			Name:        ex.Name,
			Description: *ex.Description,
			Difficulty:  ex.Difficulty,
			Position:    ex.Position,
			CreatedAt:   ex.CreatedAt,
			UpdatedAt:   ex.UpdatedAt,
			DeletedAt:   ex.DeletedAt,
			PurgeAt:     ex.PurgeAt,
		},
	}
}

func (b *ViewBuilder) Build() *view.View {
	return b.view
}

func (b *ViewBuilder) WithAliases(aliases []string) *ViewBuilder {
	b.view.Alias = aliases
	return b
}

func (b *ViewBuilder) WithEquipment(equipment view.Equipment) *ViewBuilder {
	b.view.Equipment = equipment
	return b
}

func (b *ViewBuilder) WithMuscles(muscles []view.MuscleGroup) *ViewBuilder {
	b.view.MuscleGroup = muscles
	return b
}

func (b *ViewBuilder) WithCategory(categories []view.Category) *ViewBuilder {
	b.view.Category = categories
	return b
}

func (cmd *BuildViewCommand) BuildEquipment(ctx context.Context, equipmentID *int) (view.Equipment, error) {
	// Get the equipment info
	if equipmentID == nil {
		return view.Equipment{}, nil // not an error, exercise just doesn't use an equipment
	}
	e, err := read.Equipment.GetByID(ctx, *equipmentID)
	if err != nil {
		return view.Equipment{}, fmt.Errorf("failed to read equipment %d: %w", equipmentID, err)
	}
	exAttachments, err := read.Exercise.Attachment.GetByExerciseID(ctx, cmd.ExerciseID)
	if err != nil {
		return view.Equipment{}, fmt.Errorf("failed to read exercise attachments: %w", err)
	}

	var attachments []*view.Attachment
	if len(exAttachments) != 0 {
		for _, i := range exAttachments {
			a, err := read.Attachment.GetByID(ctx, *i.AttachmentID)
			if err != nil {
				return view.Equipment{}, fmt.Errorf("failed to read attachment: %w", err)
			}

			attachment := &view.Attachment{
				Name: a.Name,
				Type: a.Type,
			}

			attachments = append(attachments, attachment)
		}
	}

	eq := view.Equipment{
		Name:        e.Name,
		Description: *e.Description,
		Type:        e.Type,
		Attachment:  attachments,
	}

	return eq, nil
}

func (cmd *BuildViewCommand) BuildMuscle(ctx context.Context) ([]view.MuscleGroup, error) {
	muscles, err := read.Exercise.Muscle.GetByExerciseID(ctx, cmd.ExerciseID)
	if err != nil {
		return nil, fmt.Errorf("failed to read exercise muscles: %w", err)
	}

	muscleView := []view.MuscleGroup{}
	if len(muscles) == 0 {
		return muscleView, nil
	}

	seen := make(map[string][]view.Muscle)

	for _, i := range muscles {
		m, err := read.Muscle.GetByID(ctx, i.MuscleID)
		if err != nil {
			return nil, fmt.Errorf("failed to read muscle: %w", err)
		}
		seen[m.Group] = append(seen[m.Group], view.Muscle{
			Name:       m.Name,
			Activation: i.Activation,
		})
	}

	for group, muscle := range seen {
		muscleView = append(muscleView, view.MuscleGroup{
			Group:  group,
			Muscle: muscle,
		})
	}
	return muscleView, nil
}

func (cmd *BuildViewCommand) BuildCategory(ctx context.Context) ([]view.Category, error) {
	categories, err := read.Exercise.Category.GetByExerciseID(ctx, cmd.ExerciseID)
	if err != nil {
		return nil, fmt.Errorf("failed to read exercise categories: %w", err)
	}

	categoryView := []view.Category{}
	if len(categories) == 0 {
		return categoryView, nil
	}

	seen := make(map[string][]string)

	for _, i := range categories {
		c, err := read.Category.GetByID(ctx, i.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to read category: %w", err)
		}
		seen[c.Type] = append(seen[c.Type], c.Name)
	}

	for categoryType, categories := range seen {
		categoryView = append(categoryView, view.Category{
			Type: categoryType,
			Name: categories,
		})
	}
	return categoryView, nil
}

func (cmd *BuildViewCommand) BuildAlias(ctx context.Context) ([]string, error) {
	aliases, err := read.Exercise.Alias.GetByExerciseID(ctx, cmd.ExerciseID)
	if err != nil {
		return nil, fmt.Errorf("failed to read exercise aliases: %w", err)
	}

	aliasView := []string{}
	if len(aliases) == 0 {
		return aliasView, nil
	}

	for _, i := range aliases {
		aliasView = append(aliasView, i.Name)
	}

	return aliasView, nil
}

func init() {
	register(&BuildViewCommand{})
}
