package commands

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
)

type UpdateMuscleCommand struct {
	ID        int    `json:"id"`
	GroupType string `json:"group_type"`
	Name      string `json:"name"`
}

type UpdateMuscleResp struct{}

func (cmd *UpdateMuscleCommand) Handle(ctx context.Context) (any, error) {
	existing, err := read.Muscle.GetByID(ctx, cmd.ID)
	if err != nil {
		return UpdateMuscleResp{}, fmt.Errorf("failed to read muscle: %w", err)
	}

	if cmd.Name != "" {
		existing.Name = cmd.Name
	}
	if cmd.GroupType != "" && cmd.GroupType != existing.Group.ToString() {
		groupType, err := muscle.NewMuscleGroupType(cmd.GroupType)
		if err != nil {
			return nil, err
		}
		existing.Group = groupType
	}
	existing.Touch()

	err = write.Muscle.Update(ctx, *existing)
	if err != nil {
		return UpdateMuscleResp{}, fmt.Errorf("failed to update muscle: %w", err)
	}

	return UpdateMuscleResp{}, nil
}

func init() {
	register(&UpdateMuscleCommand{})
}
