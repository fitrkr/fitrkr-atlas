package commands

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
)

type MuscleReq struct {
	MuscleID   int    `json:"muscle_id"`
	Activation string `json:"activation"`
}

type UpdateExerciseMuscleCommand struct {
	Remove     []int       `json:"remove,omitempty"`
	Add        []MuscleReq `json:"add,omitempty"`
	ExerciseID int         `json:"exercise_id"`
}

type UpdateExerciseMuscleResp struct{}

func (cmd *UpdateExerciseMuscleCommand) Handle(ctx context.Context) (any, error) {
	if len(cmd.Remove) != 0 {
		err := write.Exercise.Muscle.Delete(ctx, cmd.Remove)
		if err != nil {
			return nil, fmt.Errorf("failed to delete exercise muscles: %w", err)
		}
	}

	if len(cmd.Add) != 0 {
		var muscles []exercise.ExerciseMuscle
		for _, m := range cmd.Add {
			level, err := exercise.NewActivationLevel(m.Activation)
			if err != nil {
				return nil, fmt.Errorf("failed to create activation level: %w", err)
			}

			muscle, err := exercise.NewExerciseMuscle(cmd.ExerciseID, m.MuscleID, level.ToString())
			if err != nil {
				return nil, fmt.Errorf("failed to create exercise muscle: %w", err)
			}

			muscles = append(muscles, muscle)
		}

		err := write.Exercise.Muscle.Add(ctx, muscles)
		if err != nil {
			return nil, fmt.Errorf("failed to insert exercise muscles: %w", err)
		}
	}

	return UpdateExerciseMuscleResp{}, nil
}

func init() {
	register(&UpdateExerciseMuscleCommand{})
}
