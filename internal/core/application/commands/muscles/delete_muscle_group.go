package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteMuscleGroupCommand struct {
	ID    int `json:"id"`
	Write ports.MuscleGroupWrite
}

type DeleteMuscleGroupResp struct{}

func (cmd *DeleteMuscleGroupCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrMuscleGroupNotFound {
			logr.Get().Error("muscle group not found")
			return DeleteMuscleGroupResp{}, ports.ErrMuscleGroupNotFound
		}
		logr.Get().Errorf("failed to delete muscle group: %v", err)
		return DeleteMuscleGroupResp{}, fmt.Errorf("failed to delete muscle group: %w", err)
	}

	logr.Get().Info("Muscle group deleted successfully")
	return DeleteMuscleGroupResp{}, nil
}
