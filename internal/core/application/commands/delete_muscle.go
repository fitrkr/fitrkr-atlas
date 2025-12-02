package commands

import (
	"context"
	"fmt"

	"github.com/fitrkr/atlas/internal/core/ports"
)

type DeleteMuscleCommand struct {
	ID int `json:"id"`
}

type DeleteMuscleResp struct{}

func (cmd *DeleteMuscleCommand) Handle(ctx context.Context) (any, error) {
	err := write.Muscle.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrMuscleNotFound {
			return DeleteMuscleResp{}, ports.ErrMuscleNotFound
		}
		return DeleteMuscleResp{}, fmt.Errorf("failed to delete muscle: %w", err)
	}

	return DeleteMuscleResp{}, nil
}

func init() {
	register(&DeleteMuscleCommand{})
}
