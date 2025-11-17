package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
)

type GetMuscleByIDQuery struct {
	ID int `json:"id"`
}

type GetMuscleByIDResp struct {
	Muscle *muscle.Muscle
}

func (qry *GetMuscleByIDQuery) Handle(ctx context.Context) (any, error) {
	m, err := read.Muscle.GetByID(ctx, qry.ID)
	if err != nil {
		return GetMuscleByIDResp{}, fmt.Errorf("failed to read muscle: %w", err)
	}

	return GetMuscleByIDResp{Muscle: m}, nil
}

func init() {
	register(&GetMuscleByIDQuery{})
}
