package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateMuscleGroupCommand struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Write       ports.MuscleGroupWrite
	Read        ports.MuscleGroupRead
}

type UpdateMuscleGroupResp struct{}

func (cmd *UpdateMuscleGroupCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.GetByID(ctx, cmd.ID)
	if err != nil {
		logr.Get().Errorf("failed to get muscle group: %v", err)
		return UpdateMuscleGroupResp{}, fmt.Errorf("failed to get muscle group: %w", err)
	}

	if cmd.Name != "" {
		name, err := muscle.NewMuscleGroupType(cmd.Name)
		if err != nil {
			logr.Get().Errorf("failed to create new muscle group type: %v", err)
			return UpdateMuscleGroupResp{}, fmt.Errorf("failed to create new muscle group type: %w", err)
		}
		existing.Name = name
	}
	if cmd.Description != "" {
		existing.Description = &cmd.Description
	}

	existing.Touch()

	err = cmd.Write.Update(ctx, *existing)
	if err != nil {
		logr.Get().Errorf("failed to update muscle group: %v", err)
		return UpdateMuscleGroupResp{}, fmt.Errorf("failed to update muscle group: %w", err)
	}

	return UpdateMuscleGroupResp{}, nil
}
