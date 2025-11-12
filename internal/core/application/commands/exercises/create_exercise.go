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

type CreateExerciseCommand struct {
	Name         string                      `json:"name"`
	Description  string                      `json:"description"`
	Difficulty   string                      `json:"difficulty"`
	BodyPosition string                      `json:"body_position"`
	Type         string                      `json:"type"`
	Equipment    CreateExerciseEquipmentReq  `json:"equipment"`
	Muscle       []CreateExerciseMuscleReq   `json:"muscle"`
	Category     []CreateExerciseCategoryReq `json:"category"`

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
	muscleGroupView, err := cmd.addMuscleGroup(ctx, *exercise.ID)
	if err != nil {
		return nil, err
	}

	// Add categories if any
	categoryView, err := cmd.addCategory(ctx, *exercise.ID)
	if err != nil {
		return nil, err
	}

	err = cmd.addView(ctx, exercise, equipmentView, muscleGroupView, categoryView)
	if err != nil {
		return nil, err
	}

	// Add it to denormalized table
	return CreateExerciseResp{ID: *exercise.ID}, nil
}
