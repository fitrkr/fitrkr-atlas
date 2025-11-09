package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteMuscleGroupCommand struct {
	ID    int `json:"id"`
	Write ports.Write
}

type DeleteMuscleGroupResp struct{}

func (cmd *DeleteMuscleGroupCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Muscle.Group.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrMuscleGroupNotFound {
			return DeleteMuscleGroupResp{}, ports.ErrMuscleGroupNotFound
		}
		return DeleteMuscleGroupResp{}, fmt.Errorf("failed to delete muscle group: %w", err)
	}

	return DeleteMuscleGroupResp{}, nil
}
