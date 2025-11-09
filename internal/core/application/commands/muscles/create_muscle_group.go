package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateMuscleGroupCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Write       ports.Write
}

type CreateMuscleGroupResp struct{}

func (cmd *CreateMuscleGroupCommand) Handle(ctx context.Context) (any, error) {
	name, err := muscle.NewMuscleGroupType(cmd.Name)
	if err != nil {
		return CreateMuscleGroupResp{}, fmt.Errorf("failed to create new muscle group type: %w", err)
	}

	m, err := muscle.NewGroup(name, cmd.Description)
	if err != nil {
		return CreateMuscleGroupResp{}, fmt.Errorf("failed to create new muscle group: %w", err)
	}

	err = cmd.Write.Muscle.Group.Add(ctx, m)
	if err != nil {
		return CreateMuscleGroupResp{}, fmt.Errorf("failed to add muscle group: %w", err)
	}

	return CreateMuscleGroupResp{}, nil
}
