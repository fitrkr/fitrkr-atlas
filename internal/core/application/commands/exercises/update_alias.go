package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateAlias struct {
	Alias  exercise.Alias
	Action Action
}

type UpdateAliasCommand struct {
	Aliases    []UpdateAlias
	ExerciseID int
	Write      ports.Write
	Read       ports.Read
}

type UpdateAliasResp struct{}

func (cmd *UpdateAliasCommand) Handle(ctx context.Context) (any, error) {
	for _, i := range cmd.Aliases {
		switch i.Action {
		case Add:
			err := cmd.Write.Exercise.Alias.Add(ctx, i.Alias)
			if err != nil {
				return nil, fmt.Errorf("failed to insert alias: %w", err)
			}
		case Remove:
			if i.Alias.ID == nil {
				return nil, fmt.Errorf("invalid alias: missing ID for removal")
			}
			err := cmd.Write.Exercise.Alias.Delete(ctx, *i.Alias.ID)
			if err != nil {
				if err == ports.ErrExerciseAliasNotFound {
					return UpdateAliasResp{}, ports.ErrExerciseAliasNotFound
				}
				return UpdateAliasResp{}, fmt.Errorf("failed to remove exercise alias: %w", err)
			}
		default:
			return nil, fmt.Errorf("unsupported action: %d", i.Action)
		}
	}

	vb := &ViewBuilder{Write: cmd.Write, Read: cmd.Read}
	err := vb.RebuildView(ctx, cmd.ExerciseID)
	if err != nil {
		return nil, fmt.Errorf("failed to rebuild view for exercise %d: %w", cmd.ExerciseID)
	}

	return UpdateAliasResp{}, nil
}
