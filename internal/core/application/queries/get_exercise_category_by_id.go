package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/exercise"
)

type GetExerciseCategoryByIDQuery struct {
	ID int `json:"id"`
}

type GetExerciseCategoryByIDResp struct {
	Category *exercise.ExerciseCategory
}

func (qry *GetExerciseCategoryByIDQuery) Handle(ctx context.Context) (any, error) {
	category, err := read.Exercise.Category.GetByID(ctx, qry.ID)
	if err != nil {
		return GetExerciseCategoryByIDResp{}, fmt.Errorf("failed to get exercise category: %w", err)
	}

	return GetExerciseCategoryByIDResp{Category: category}, nil
}

func init() {
	register(&GetExerciseCategoryByIDQuery{})
}
