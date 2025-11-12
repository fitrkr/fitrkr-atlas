package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/view"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type Action int

const (
	Add Action = iota
	Remove
)

// Helper functions
func (cmd *CreateExerciseCommand) addEquipment(ctx context.Context, exerciseID int) (*view.Equipment, error) {
	attachment := view.Attachment{}

	e, err := cmd.Read.Equipment.GetByID(ctx, cmd.Equipment.EquipmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to read equipment %d: %w", cmd.Equipment.EquipmentID, err)
	}

	if cmd.Equipment.AttachmentID != nil {
		// If there is an attachment, validate it
		a, err := cmd.Read.Attachment.GetByID(ctx, *cmd.Equipment.AttachmentID)
		if err != nil {
			return nil, fmt.Errorf("failed to read attachment %d: %w", cmd.Equipment.AttachmentID, err)
		}
		attachment.ID = *cmd.Equipment.AttachmentID
		attachment.Name = a.Name
		attachment.Type = a.Type.ToString()
	}

	exerciseEquipment, err := exercise.NewExerciseEquipment(exerciseID, cmd.Equipment.EquipmentID, cmd.Equipment.AttachmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to create new exercise equipment: %w", err)
	}
	err = cmd.Write.Exercise.Equipment.Add(ctx, exerciseEquipment)
	if err != nil {
		return nil, fmt.Errorf("failed to insert exercise equipment: %w", err)
	}

	equipmentView := view.Equipment{
		ID:          *e.ID,
		Name:        e.Name,
		Description: *e.Description,
		Type:        e.Type.ToString(),
		Attachment:  attachment,
	}

	return &equipmentView, nil
}

func (cmd *CreateExerciseCommand) addMuscle(ctx context.Context, exerciseID int) ([]view.MuscleGroup, error) {
	muscleView := []view.MuscleGroup{}
	if len(cmd.Muscle) == 0 {
		return muscleView, nil
	}

	seen := make(map[int][]view.Muscle)

	for _, i := range cmd.Muscle {
		m, err := cmd.Read.Muscle.GetByID(ctx, i.MuscleID)
		if err != nil {
			return nil, ports.ErrMuscleNotFound
		}
		level, err := exercise.NewActivationLevel(i.ActivationLevel)
		if err != nil {
			return nil, fmt.Errorf("failed to create new activation level: %w", err)
		}

		muscle, err := exercise.NewExerciseMuscle(exerciseID, i.MuscleID, level)
		if err != nil {
			return nil, fmt.Errorf("failed to create new exercise muscle: %w", err)
		}

		err = cmd.Write.Exercise.Muscle.Add(ctx, muscle)
		if err != nil {
			return nil, fmt.Errorf("failed to add new exercise muscle: %w", err)
		}

		seen[m.MuscleGroupID] = append(seen[m.MuscleGroupID], view.Muscle{
			ID:         *m.ID,
			Name:       m.Name,
			Activation: level.ToString(),
		})
	}

	for muscleGroupID, muscle := range seen {
		mg, err := cmd.Read.Muscle.Group.GetByID(ctx, muscleGroupID)
		if err != nil {
			return nil, fmt.Errorf("failed to read muscle group: %w", err)
		}

		// Add to view grouping
		muscleView = append(muscleView, view.MuscleGroup{
			ID:          muscleGroupID,
			Name:        string(mg.Name),
			Description: *mg.Description,
			Muscle:      muscle,
		})
	}
	return muscleView, nil
}

func (cmd *CreateExerciseCommand) addCategory(ctx context.Context, exerciseID int) ([]view.Category, error) {
	categoryView := []view.Category{}
	if len(cmd.Category) == 0 {
		return categoryView, nil
	}

	seen := make(map[int][]view.Subcategory)

	for _, i := range cmd.Category {
		sc, err := cmd.Read.Category.Subcategory.GetByID(ctx, i.SubcategoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to read subcategory: %w", err)
		}
		category, err := exercise.NewExerciseCategory(exerciseID, i.SubcategoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to create new exercise category: %w", err)
		}
		err = cmd.Write.Exercise.Category.Add(ctx, category)
		if err != nil {
			return nil, fmt.Errorf("failed to add new exercise category: %w", err)
		}

		// Add to view grouping
		seen[sc.CategoryID] = append(seen[sc.CategoryID], view.Subcategory{
			ID:   *sc.ID,
			Name: sc.Name,
		})
	}

	for categoryID, subcategories := range seen {
		c, err := cmd.Read.Category.GetByID(ctx, categoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to read category: %w", err)
		}

		categoryView = append(categoryView, view.Category{
			ID:          categoryID,
			Name:        string(c.Name),
			Subcategory: subcategories,
		})
	}
	return categoryView, nil
}

func (cmd *CreateExerciseCommand) addView(ctx context.Context, ex exercise.Exercise, equipmentView view.Equipment, muscleGroupView []view.MuscleGroup, categoryView []view.Category) error {
	v := view.NewExerciseView(
		*ex.ID,
		ex.Name,
		*ex.Description,
		ex.Type.ToString(),
		ex.Difficulty.ToString(),
		string(ex.BodyPosition),
		muscleGroupView,
		equipmentView,
		categoryView,
		ex.CreatedAt,
		ex.UpdatedAt,
		ex.DeletedAt,
		ex.PurgedAt,
	)

	_, err := cmd.Write.View.Add(ctx, v)
	if err != nil {
		return fmt.Errorf("failed to insert view: %w", err)
	}

	return nil
}
