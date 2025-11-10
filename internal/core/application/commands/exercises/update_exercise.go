package exercises

import (
	"context"
	"fmt"
	"strings"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateExerciseCommand struct {
	ID           int
	Name         string
	Description  string
	Difficulty   string
	BodyPosition string
	Write        ports.Write
	Read         ports.Read
}

type UpdateExerciseResp struct {
	Exercise exercise.Exercise
}

func (cmd *UpdateExerciseCommand) Handle(ctx context.Context) (any, error) {
	exists, err := cmd.Read.Exercise.GetByID(ctx, cmd.ID)
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

	if cmd.BodyPosition != "" {
		bodyPosition, err := exercise.NewBodyPosition(cmd.BodyPosition)
		if err != nil {
			return CreateExerciseResp{}, exercise.ErrInvalidBodyPosition
		}
		exists.BodyPosition = bodyPosition
	}

	exists.Touch()

	err = cmd.Write.Exercise.Update(ctx, *exists)
	if err != nil {
		return nil, fmt.Errorf("failed to update exercise: %w", err)
	}

	return UpdateExerciseResp{}, nil
}
