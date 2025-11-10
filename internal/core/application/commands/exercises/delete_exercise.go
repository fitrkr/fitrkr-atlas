package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteExerciseCommand struct {
	ID    int
	Write ports.Write
	Read  ports.Read
}

type DeleteExerciseResp struct{}

func (cmd *DeleteExerciseCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Exercise.Delete(ctx, cmd.ID)
	// TODO make sure to delete all the other junction tables
	if err != nil {
		if err == ports.ErrExerciseNotFound {
			return DeleteExerciseResp{}, ports.ErrExerciseNotFound
		}
		return DeleteExerciseResp{}, fmt.Errorf("failed to exercise muscle: %w", err)
	}

	return DeleteExerciseResp{}, nil
}
