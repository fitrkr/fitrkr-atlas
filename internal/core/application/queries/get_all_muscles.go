package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
)

type GetAllMusclesQuery struct{}

type GetAllMusclesResp struct {
	Muscles []*muscle.Muscle
}

func (qry *GetAllMusclesQuery) Handle(ctx context.Context) (any, error) {
	muscles, err := read.Muscle.GetAll(ctx)
	if err != nil {
		return GetAllMusclesResp{}, fmt.Errorf("failed to read muscles: %w", err)
	}

	return GetAllMusclesResp{Muscles: muscles}, nil
}

func init() {
	register(&GetAllMusclesQuery{})
}
