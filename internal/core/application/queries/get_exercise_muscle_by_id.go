package queries

import (
	"context"
	"fmt"

	"github.com/fitrkr/atlas/internal/core/domain/exercise"
)

type GetExerciseMuscleByIDQuery struct {
	ID int `json:"id"`
}

type GetExerciseMuscleByIDResp struct {
	Muscle *exercise.ExerciseMuscle
}

func (qry *GetExerciseMuscleByIDQuery) Handle(ctx context.Context) (any, error) {
	muscle, err := read.Exercise.Muscle.GetByID(ctx, qry.ID)
	if err != nil {
		return GetExerciseMuscleByIDResp{}, fmt.Errorf("failed to get exercise muscle: %w", err)
	}

	return GetExerciseMuscleByIDResp{Muscle: muscle}, nil
}

func init() {
	register(&GetExerciseMuscleByIDQuery{})
}
