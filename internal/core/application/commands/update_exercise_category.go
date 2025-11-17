package commands

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
)

type UpdateExerciseCategoryCommand struct {
	Remove     []int
	Add        []int // category_ids
	ExerciseID int
}

type UpdateExerciseCategoryResp struct{}

func (cmd *UpdateExerciseCategoryCommand) Handle(ctx context.Context) (any, error) {
	if len(cmd.Remove) != 0 {
		err := write.Exercise.Category.Delete(ctx, cmd.Remove)
		if err != nil {
			return nil, fmt.Errorf("failed to delete exercise categories: %w", err)
		}
	}

	if len(cmd.Add) != 0 {
		var categories []exercise.ExerciseCategory
		for _, categoryID := range cmd.Add {
			category, err := exercise.NewExerciseCategory(cmd.ExerciseID, categoryID)
			if err != nil {
				return nil, fmt.Errorf("failed to create exercise category: %w", err)
			}

			categories = append(categories, category)
		}

		err := write.Exercise.Category.Add(ctx, categories)
		if err != nil {
			return nil, fmt.Errorf("failed to insert exercise categories: %w", err)
		}
	}

	return UpdateExerciseCategoryResp{}, nil
}

func init() {
	register(&UpdateExerciseCategoryCommand{})
}
