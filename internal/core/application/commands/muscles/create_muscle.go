// Package muscles
package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateMuscleCommand struct {
	Name          string `json:"name"`
	MuscleGroupID int    `json:"muscle_group_id"`
	Write         ports.Write
	Read          ports.Read
}

type CreateMuscleResp struct{}

func (cmd *CreateMuscleCommand) Handle(ctx context.Context) (any, error) {
	_, err := cmd.Read.Muscle.Group.GetByID(ctx, cmd.MuscleGroupID)
	if err != nil {
		return CreateMuscleResp{}, fmt.Errorf("failed to get muscle group: %w", err)
	}

	m, err := muscle.New(cmd.Name, cmd.MuscleGroupID)
	if err != nil {
		return CreateMuscleResp{}, fmt.Errorf("failed to create new muscle: %w", err)
	}

	err = cmd.Write.Muscle.Add(ctx, m)
	if err != nil {
		return CreateMuscleResp{}, fmt.Errorf("failed to add muscle: %w", err)
	}

	return CreateMuscleResp{}, nil
}
