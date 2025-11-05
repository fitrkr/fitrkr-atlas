package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteMuscleCommand struct {
	ID    int `json:"id"`
	Write ports.MuscleWrite
}

type DeleteMuscleResp struct{}

func (cmd *DeleteMuscleCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrMuscleNotFound {
			logr.Get().Error("muscle not found")
			return DeleteMuscleResp{}, ports.ErrMuscleNotFound
		}
		logr.Get().Errorf("failed to delete muscle: %v", err)
		return DeleteMuscleResp{}, fmt.Errorf("failed to delete muscle: %w", err)
	}

	logr.Get().Info("Muscle deleted successfully")
	return DeleteMuscleResp{}, nil
}
