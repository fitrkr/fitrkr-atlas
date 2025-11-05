package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateExerciseCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
	Write       ports.ExerciseWrite
}

type CreateExerciseResp struct{}

func (cmd *CreateExerciseCommand) Handle(ctx context.Context) (any, error) {
	difficulty, err := exercise.NewDifficulty(cmd.Difficulty)
	if err != nil {
		logr.Get().Error("invalid difficulty")
		return CreateExerciseResp{}, exercise.ErrInvalidDifficulty
	}

	ex, err := exercise.New(cmd.Name, cmd.Description, difficulty)
	if err != nil {
		logr.Get().Errorf("failed to create new exercise: %v", err)
		return CreateExerciseResp{}, fmt.Errorf("failed to create new exercise: %w", err)
	}

	err = cmd.Write.Add(ctx, ex)
	if err != nil {
		logr.Get().Errorf("failed to add exercise to db: %v", err)
		return CreateExerciseResp{}, fmt.Errorf("failed to add exercise to db: %w", err)
	}

	return CreateExerciseResp{}, nil
}
