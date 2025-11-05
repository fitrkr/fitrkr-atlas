package exercises

import (
	"context"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteExerciseCommand struct {
	ID    int
	Write ports.ExerciseWrite
	Read  ports.ExerciseRead
}

type DeleteExerciseResp struct{}

func (cmd *DeleteExerciseCommand) Handle(ctx context.Context) (any, error) {
	return DeleteExerciseResp{}, nil
}
