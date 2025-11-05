package exercises

import (
	"context"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateExerciseCommand struct {
	ID          int
	Name        string
	Description string
	Difficulty  exercise.Difficulty
	Write       ports.ExerciseWrite
	Read        ports.ExerciseRead
}

type UpdateExerciseResp struct{}

func (cmd *UpdateExerciseCommand) Handle(ctx context.Context) (any, error) {
	return UpdateExerciseResp{}, nil
}
