// Package muscles
package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateMuscleCommand struct {
	Name      string `json:"name"`
	GroupType string `json:"group_type"`
	Write     ports.Write
	Read      ports.Read
}

type CreateMuscleResp struct{}

func (cmd *CreateMuscleCommand) Handle(ctx context.Context) (any, error) {
	groupType, err := muscle.NewMuscleGroupType(cmd.GroupType)
	if err != nil {
		return nil, err
	}

	m, err := muscle.New(cmd.Name, groupType)
	if err != nil {
		return CreateMuscleResp{}, fmt.Errorf("failed to create new muscle: %w", err)
	}

	err = cmd.Write.Muscle.Add(ctx, m)
	if err != nil {
		return CreateMuscleResp{}, fmt.Errorf("failed to insert muscle: %w", err)
	}

	return CreateMuscleResp{}, nil
}
