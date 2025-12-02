package commands

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type DeleteExerciseCommand struct {
	ID int `json:"id"`
}

type DeleteExerciseResp struct{}

func (cmd *DeleteExerciseCommand) Handle(ctx context.Context) (any, error) {
	err := write.Exercise.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrExerciseNotFound {
			return DeleteExerciseResp{}, ports.ErrExerciseNotFound
		}
		return DeleteExerciseResp{}, fmt.Errorf("failed to delete exercise: %w", err)
	}

	return DeleteExerciseResp{}, nil
}

func init() {
	register(&DeleteExerciseCommand{})
}
