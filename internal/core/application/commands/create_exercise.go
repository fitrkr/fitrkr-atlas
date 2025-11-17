package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateExerciseCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	Position    string `json:"position"`
	Type        string `json:"type"`
	EquipmentID *int   `json:"equipment_id"`
}

type CreateExerciseResp struct {
	ID int
}

func (cmd *CreateExerciseCommand) Handle(ctx context.Context) (any, error) {
	// Check if exercise already exists
	name := strings.TrimSpace(strings.ToLower(cmd.Name))
	if _, err := read.Exercise.GetByName(ctx, name); err == nil {
		return nil, ports.ErrDuplicateExercise
	}

	// Validate difficulty
	difficulty, err := exercise.NewDifficulty(cmd.Difficulty)
	if err != nil {
		return CreateExerciseResp{}, exercise.ErrInvalidDifficulty
	}

	// Validate body position
	position, err := exercise.NewBodyPosition(cmd.Position)
	if err != nil {
		return CreateExerciseResp{}, exercise.ErrInvalidPosition
	}

	// Create Exercise
	ex, err := exercise.New(name, cmd.Description, difficulty, position)
	if err != nil {
		return CreateExerciseResp{}, fmt.Errorf("failed to create new exercise: %w", err)
	}

	// Add exercise to DB and get exercise back
	ex, err = write.Exercise.Add(ctx, ex)
	if err != nil {
		return CreateExerciseResp{}, fmt.Errorf("failed to insert exercise: %w", err)
	}

	// Validate equipment if provided
	if cmd.EquipmentID != nil {
		equipment, err := read.Equipment.GetByID(ctx, *cmd.EquipmentID)
		if err != nil {
			return CreateExerciseResp{}, fmt.Errorf("failed to read equipment %d: %w", *cmd.EquipmentID, err)
		}
		ex.EquipmentID = equipment.ID
	}

	return CreateExerciseResp{ID: *ex.ID}, nil
}

func init() {
	register(&CreateExerciseCommand{})
}
