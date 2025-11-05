package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateMuscleCommand struct {
	ID            int    `json:"id"`
	MuscleGroupID int    `json:"muscle_group_id"`
	Name          string `json:"name"`
	Write         ports.MuscleWrite
	Read          ports.MuscleRead
	ReadGroup     ports.MuscleGroupRead
}

type UpdateMuscleResp struct{}

func (cmd *UpdateMuscleCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.GetByID(ctx, cmd.ID)
	if err != nil {
		logr.Get().Errorf("failed to get muscle: %v", err)
		return UpdateMuscleResp{}, fmt.Errorf("failed to get muscle: %w", err)
	}

	if cmd.Name != "" {
		existing.Name = cmd.Name
	}
	if cmd.MuscleGroupID > 0 {
		_, err := cmd.ReadGroup.GetByID(ctx, cmd.MuscleGroupID)
		if err != nil {
			logr.Get().Errorf("failed to get muscle group: %v", err)
			return UpdateMuscleResp{}, fmt.Errorf("failed to get muscle group: %w", err)
		}
		existing.MuscleGroupID = cmd.MuscleGroupID
	}
	existing.Touch()

	err = cmd.Write.Update(ctx, *existing)
	if err != nil {
		logr.Get().Errorf("failed to update muscle: %v", err)
		return UpdateMuscleResp{}, fmt.Errorf("failed to update muscle: %w", err)
	}

	return UpdateMuscleResp{}, nil
}
