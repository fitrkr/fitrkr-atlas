package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetExerciseByNameQuery struct {
	Name string `json:"name"`
	Read ports.Read
}

type GetExerciseByNameResp struct {
	Exercise *exercise.Exercise
}

func (qry *GetExerciseByNameQuery) Handle(ctx context.Context) (any, error) {
	ex, err := qry.Read.Exercise.GetByName(ctx, qry.Name)
	if err != nil {
		return GetExerciseByNameResp{}, fmt.Errorf("failed to get equipment: %w", err)
	}

	return GetExerciseByNameResp{Exercise: ex}, nil
}
