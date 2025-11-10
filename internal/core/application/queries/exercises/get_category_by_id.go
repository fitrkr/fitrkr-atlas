// Package exercises
package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetCategoryByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetCategoryByIDResp struct {
	Category *exercise.ExerciseCategory
}

func (qry *GetCategoryByIDQuery) Handle(ctx context.Context) (any, error) {
	category, err := qry.Read.Exercise.Category.GetByID(ctx, qry.ID)
	if err != nil {
		return GetCategoryByIDResp{}, fmt.Errorf("failed to get exercise category: %w", err)
	}

	return GetCategoryByIDResp{Category: category}, nil
}
