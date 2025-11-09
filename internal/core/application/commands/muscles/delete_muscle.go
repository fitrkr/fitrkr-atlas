package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteMuscleCommand struct {
	ID    int `json:"id"`
	Write ports.Write
}

type DeleteMuscleResp struct{}

func (cmd *DeleteMuscleCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Muscle.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrMuscleNotFound {
			return DeleteMuscleResp{}, ports.ErrMuscleNotFound
		}
		return DeleteMuscleResp{}, fmt.Errorf("failed to delete muscle: %w", err)
	}

	return DeleteMuscleResp{}, nil
}
