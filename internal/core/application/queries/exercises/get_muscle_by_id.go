package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetMuscleByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetMuscleByIDResp struct {
	Muscle *exercise.ExerciseMuscle
}

func (qry *GetMuscleByIDQuery) Handle(ctx context.Context) (any, error) {
	muscle, err := qry.Read.Exercise.Muscle.GetByID(ctx, qry.ID)
	if err != nil {
		return GetMuscleByIDResp{}, fmt.Errorf("failed to get exercise muscle: %w", err)
	}

	return GetMuscleByIDResp{Muscle: muscle}, nil
}
