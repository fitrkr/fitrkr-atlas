package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateExerciseCommand struct {
	ID          int
	Name        string
	Description string
	Difficulty  string
	Position    string
	EquipmentID *int
}

type UpdateExerciseResp struct {
	Exercise exercise.Exercise
}

func (cmd *UpdateExerciseCommand) Handle(ctx context.Context) (any, error) {
	exists, err := read.Exercise.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, ports.ErrExerciseNotFound
	}

	if cmd.Name != "" {
		name := strings.ToLower(cmd.Name)
		exists.Name = strings.TrimSpace(name)
	}

	if cmd.Description != "" {
		description := strings.ToLower(cmd.Description)
		exists.Name = strings.TrimSpace(description)
	}

	if cmd.Difficulty != "" {
		difficulty, err := exercise.NewDifficulty(cmd.Difficulty)
		if err != nil {
			return UpdateExerciseResp{}, exercise.ErrInvalidDifficulty
		}
		exists.Difficulty = difficulty
	}

	if cmd.Position != "" {
		position, err := exercise.NewBodyPosition(cmd.Position)
		if err != nil {
			return CreateExerciseResp{}, exercise.ErrInvalidPosition
		}
		exists.Position = position
	}

	// Validate equipment if provided
	if cmd.EquipmentID != nil {
		equipment, err := read.Equipment.GetByID(ctx, *cmd.EquipmentID)
		if err != nil {
			return CreateExerciseResp{}, fmt.Errorf("failed to read equipment %d: %w", *cmd.EquipmentID, err)
		}
		exists.EquipmentID = equipment.ID
	}

	exists.Touch()

	_, err = write.Exercise.Update(ctx, *exists)
	if err != nil {
		return nil, fmt.Errorf("failed to update exercise: %w", err)
	}

	return UpdateExerciseResp{}, nil
}

func init() {
	register(&UpdateExerciseCommand{})
}
