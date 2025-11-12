// Package exercises
package exercises

import (
	"context"
	"fmt"
	"strings"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/view"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateExerciseEquipmentReq struct {
	EquipmentID  int  `json:"equipment_id"`
	AttachmentID *int `json:"attachment_id"`
}

type CreateExerciseMuscleReq struct {
	MuscleID        int    `json:"muscle_id"`
	ActivationLevel string `json:"activation_level"`
}

type CreateExerciseCategoryReq struct {
	SubcategoryID int `json:"subcategory_id"`
}

type CreateExerciseCommand struct {
	Name         string                       `json:"name"`
	Description  string                       `json:"description"`
	Difficulty   string                       `json:"difficulty"`
	BodyPosition string                       `json:"body_position"`
	Type         string                       `json:"type"`
	Equipment    []CreateExerciseEquipmentReq `json:"equipment"`
	Muscle       []CreateExerciseMuscleReq    `json:"muscle"`
	Category     []CreateExerciseCategoryReq  `json:"category"`

	Write ports.Write
	Read  ports.Read
}

type CreateExerciseResp struct {
	ID int
}

func (cmd *CreateExerciseCommand) Handle(ctx context.Context) (any, error) {
	// Check if exercise already exists
	name := strings.TrimSpace(strings.ToLower(cmd.Name))
	if _, err := cmd.Read.Exercise.GetByName(ctx, name); err == nil {
		return nil, ports.ErrDuplicateExercise
	}

	// Validate difficulty
	difficulty, err := exercise.NewDifficulty(cmd.Difficulty)
	if err != nil {
		return CreateExerciseResp{}, exercise.ErrInvalidDifficulty
	}

	// Validate body position
	bodyPosition, err := exercise.NewBodyPosition(cmd.BodyPosition)
	if err != nil {
		return CreateExerciseResp{}, exercise.ErrInvalidBodyPosition
	}

	exerciseType, err := exercise.NewExerciseType(cmd.Type)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new exercise type: %w", err)
	}

	// Create Exercise
	ex, err := exercise.New(cmd.Name, cmd.Description, difficulty, bodyPosition, exerciseType)
	if err != nil {
		return CreateExerciseResp{}, fmt.Errorf("failed to create new exercise: %w", err)
	}

	// Add exercise to DB and get id back
	exercise, err := cmd.Write.Exercise.Add(ctx, ex)
	if err != nil {
		return CreateExerciseResp{}, fmt.Errorf("failed to add exercise to db: %w", err)
	}

	// Add equipemts if any
	equipmentView, err := cmd.addEquipment(ctx, *exercise.ID)
	if err != nil {
		return nil, err
	}

	// Add muscles if any
	muscleView, err := cmd.addMuscle(ctx, *exercise.ID)
	if err != nil {
		return nil, err
	}

	// Add categories if any
	categoryView, err := cmd.addCategory(ctx, *exercise.ID)
	if err != nil {
		return nil, err
	}

	err = cmd.addView(ctx, exercise, equipmentView, muscleView, categoryView)
	if err != nil {
		return nil, err
	}

	// Add it to denormalized table
	return CreateExerciseResp{ID: *exercise.ID}, nil
}

// Helper functions
func (cmd *CreateExerciseCommand) addEquipment(ctx context.Context, exerciseID int) ([]view.Equipment, error) {
	equipmentView := []view.Equipment{}
	if len(cmd.Equipment) == 0 {
		return equipmentView, nil
	}

	seen := make(map[int][]view.Attachment)
	equipCache := make(map[int]view.Equipment)

	for _, i := range cmd.Equipment {
		_, ok := equipCache[i.EquipmentID]
		if !ok {
			e, err := cmd.Read.Equipment.GetByID(ctx, i.EquipmentID)
			if err != nil {
				return nil, fmt.Errorf("failed to get equipment %d: %w", i.EquipmentID, err)
			}

			equipCache[i.EquipmentID] = view.Equipment{
				ID:          *e.ID,
				Name:        e.Name,
				Description: *e.Description,
				Type:        e.Type.ToString(),
			}
		}

		if i.AttachmentID != nil {
			// If there is an attachment, validate it
			a, err := cmd.Read.Attachment.GetByID(ctx, *i.AttachmentID)
			if err != nil {
				return nil, fmt.Errorf("failed to get attachment: %w", err)
			}
			attachment := view.Attachment{
				ID:   *i.AttachmentID,
				Name: a.Name,
				Type: a.Type.ToString(),
			}

			seen[i.EquipmentID] = append(seen[i.EquipmentID], attachment)
		}

		exerciseEquipment, err := exercise.NewExerciseEquipment(exerciseID, i.EquipmentID, i.AttachmentID)
		if err != nil {
			return nil, fmt.Errorf("failed to create new exercise equipment: %w", err)
		}
		err = cmd.Write.Exercise.Equipment.Add(ctx, exerciseEquipment)
		if err != nil {
			return nil, fmt.Errorf("failed to add new exercise equipment: %w", err)
		}
	}

	for equipmentID, e := range equipCache {
		e.Attachment = seen[equipmentID]
		equipmentView = append(equipmentView, e)
	}

	return equipmentView, nil
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

func (cmd *CreateExerciseCommand) addView(ctx context.Context, ex exercise.Exercise, equipmentView []view.Equipment, muscleView []view.MuscleGroup, categoryView []view.Category) error {
	v := view.NewExerciseView(
		*ex.ID,
		ex.Name,
		*ex.Description,
		ex.Type.ToString(),
		ex.Difficulty.ToString(),
		string(ex.BodyPosition),
		muscleView,
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
