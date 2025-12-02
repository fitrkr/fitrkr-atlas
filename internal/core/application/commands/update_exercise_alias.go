package commands

import (
	"context"
	"fmt"

	"github.com/fitrkr/atlas/internal/core/domain/exercise"
)

type UpdateAliasCommand struct {
	Remove     []int            `json:"remove,omitempty"`
	Add        []exercise.Alias `json:"add,omitempty"`
	ExerciseID int              `json:"exercise_id"`
}

type UpdateAliasResp struct{}

func (cmd *UpdateAliasCommand) Handle(ctx context.Context) (any, error) {
	if len(cmd.Remove) != 0 {
		err := write.Exercise.Alias.Delete(ctx, cmd.Remove)
		if err != nil {
			return nil, fmt.Errorf("failed to delete exercise aliases: %w", err)
		}
	}

	if len(cmd.Add) != 0 {
		var Aliases []exercise.Alias
		for _, i := range cmd.Add {
			alias, err := exercise.NewAlias(cmd.ExerciseID, i.Name, i.LanguageCode)
			if err != nil {
				return nil, fmt.Errorf("failed to create exercise alias: %w", err)
			}

			Aliases = append(Aliases, alias)
		}

		err := write.Exercise.Alias.Add(ctx, Aliases)
		if err != nil {
			return nil, fmt.Errorf("failed to insert aliases: %w", err)
		}
	}

	return UpdateAliasResp{}, nil
}

func init() {
	register(&UpdateAliasCommand{})
}
