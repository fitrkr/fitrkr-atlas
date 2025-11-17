package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
)

type GetExerciseByNameQuery struct {
	Name string `json:"name"`
}

type GetExerciseByNameResp struct {
	Exercise *exercise.Exercise
}

func (qry *GetExerciseByNameQuery) Handle(ctx context.Context) (any, error) {
	ex, err := read.Exercise.GetByName(ctx, qry.Name)
	if err != nil {
		return GetExerciseByNameResp{}, fmt.Errorf("failed to get equipment: %w", err)
	}

	return GetExerciseByNameResp{Exercise: ex}, nil
}

func init() {
	register(&GetExerciseByNameQuery{})
}
