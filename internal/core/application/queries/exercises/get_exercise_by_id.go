package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetExerciseByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetExerciseByIDResp struct {
	Exercise *exercise.Exercise
}

func (qry *GetExerciseByIDQuery) Handle(ctx context.Context) (any, error) {
	ex, err := qry.Read.Exercise.GetByID(ctx, qry.ID)
	if err != nil {
		return GetExerciseByIDResp{}, fmt.Errorf("failed to get equipment: %w", err)
	}

	return GetExerciseByIDResp{Exercise: ex}, nil
}
