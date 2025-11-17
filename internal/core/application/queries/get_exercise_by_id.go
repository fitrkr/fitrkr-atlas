package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
)

type GetExerciseByIDQuery struct {
	ID int `json:"id"`
}

type GetExerciseByIDResp struct {
	Exercise *exercise.Exercise
}

func (qry *GetExerciseByIDQuery) Handle(ctx context.Context) (any, error) {
	ex, err := read.Exercise.GetByID(ctx, qry.ID)
	if err != nil {
		return GetExerciseByIDResp{}, fmt.Errorf("failed to get equipment: %w", err)
	}

	return GetExerciseByIDResp{Exercise: ex}, nil
}

func init() {
	register(&GetExerciseByIDQuery{})
}
