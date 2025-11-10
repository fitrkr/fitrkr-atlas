// Package exercises
package exercises

import (
	"context"
	"fmt"
	"strings"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
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

type CreateExerciseAliasReq struct {
	Name         string `json:"name"`
	LanguageCode string `json:"language_code"`
}

type CreateExerciseCommand struct {
	Name         string                       `json:"name"`
	Description  *string                      `json:"description"`
	Difficulty   string                       `json:"difficulty"`
	BodyPosition string                       `json:"body_position"`
	Type         string                       `json:"type"`
	Alias        []CreateExerciseAliasReq     `json:"alias"`
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
	_, err := cmd.Read.Exercise.GetByName(ctx, name)
	if err != nil {
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
	ex, err := exercise.New(cmd.Name, *cmd.Description, difficulty, bodyPosition, exerciseType)
	if err != nil {
		return CreateExerciseResp{}, fmt.Errorf("failed to create new exercise: %w", err)
	}

	// Add exercise to DB and get id back
	exerciseID, err := cmd.Write.Exercise.Add(ctx, ex)
	if err != nil {
		return CreateExerciseResp{}, fmt.Errorf("failed to add exercise to db: %w", err)
	}

	// Add equipemts if any
	err = cmd.addEquipment(exerciseID, ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	// Add muscles if any
	err = cmd.addMuscle(exerciseID, ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	// Add categories if any
	err = cmd.addCategory(exerciseID, ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	// Add aliases if any
	err = cmd.addAlias(exerciseID, ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return CreateExerciseResp{ID: exerciseID}, nil
}

// Helper functions

func (cmd *CreateExerciseCommand) addEquipment(exerciseID int, ctx context.Context) error {
	if len(cmd.Equipment) != 0 {
		for _, e := range cmd.Equipment {
			// Validate new values
			_, err := cmd.Read.Equipment.GetByID(ctx, e.EquipmentID)
			if err != nil {
				return ports.ErrEquipmentNotFound
			}

			exerciseEquipment, err := exercise.NewExerciseEquipment(exerciseID, e.EquipmentID, e.AttachmentID)
			if err != nil {
				return fmt.Errorf("failed to create new exercise equipment: %w", err)
			}
			err = cmd.Write.Exercise.Equipment.Add(ctx, exerciseEquipment)
			if err != nil {
				return fmt.Errorf("failed to add new exercise equipment: %w", err)
			}
		}
	}
	return nil
}

func (cmd *CreateExerciseCommand) addMuscle(exerciseID int, ctx context.Context) error {
	if len(cmd.Muscle) != 0 {
		for _, m := range cmd.Muscle {
			// Validate new values
			_, err := cmd.Read.Muscle.GetByID(ctx, m.MuscleID)
			if err != nil {
				return ports.ErrMuscleNotFound
			}

			level, err := exercise.NewActivationLevel(m.ActivationLevel)
			if err != nil {
				return fmt.Errorf("failed to create new activation level: %w", err)
			}

			muscle, err := exercise.NewExerciseMuscle(exerciseID, m.MuscleID, level)
			if err != nil {
				return fmt.Errorf("failed to create new exercise muscle: %w", err)
			}

			err = cmd.Write.Exercise.Muscle.Add(ctx, muscle)
			if err != nil {
				return fmt.Errorf("failed to add new exercise muscle: %w", err)
			}
		}
	}
	return nil
}

func (cmd *CreateExerciseCommand) addCategory(exerciseID int, ctx context.Context) error {
	if len(cmd.Category) != 0 {
		for _, c := range cmd.Category {
			category, err := exercise.NewExerciseCategory(exerciseID, c.SubcategoryID)
			if err != nil {
				return fmt.Errorf("failed to create new exercise category: %w", err)
			}
			err = cmd.Write.Exercise.Category.Add(ctx, category)
			if err != nil {
				return fmt.Errorf("failed to add new exercise category: %w", err)
			}
		}
	}
	return nil
}

func (cmd *CreateExerciseCommand) addAlias(exerciseID int, ctx context.Context) error {
	if len(cmd.Alias) != 0 {
		for _, a := range cmd.Alias {
			a, err := exercise.NewAlias(exerciseID, a.Name, a.LanguageCode)
			if err != nil {
				return fmt.Errorf("failed to create new exercise alias: %w", err)
			}
			err = cmd.Write.Exercise.Alias.Add(ctx, a)
			if err != nil {
				return fmt.Errorf("failed to add new exercise alias: %w", err)
			}
		}
	}
	return nil
}
