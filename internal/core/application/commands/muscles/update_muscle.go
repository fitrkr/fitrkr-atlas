package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateMuscleCommand struct {
	ID        int    `json:"id"`
	GroupType string `json:"group_type"`
	Name      string `json:"name"`
	Write     ports.Write
	Read      ports.Read
}

type UpdateMuscleResp struct{}

func (cmd *UpdateMuscleCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.Muscle.GetByID(ctx, cmd.ID)
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

	err = cmd.Write.Muscle.Update(ctx, *existing)
	if err != nil {
		return UpdateMuscleResp{}, fmt.Errorf("failed to update muscle: %w", err)
	}

	return UpdateMuscleResp{}, nil
}
