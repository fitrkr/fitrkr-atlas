package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/muscle"
)

type GetMusclesByGroupTypeQuery struct {
	GroupType string `json:"group_type"`
}

type GetMusclesByGroupTypeResp struct {
	Muscle []*muscle.Muscle
}

func (qry *GetMusclesByGroupTypeQuery) Handle(ctx context.Context) (any, error) {
	m, err := read.Muscle.GetByGroupType(ctx, qry.GroupType)
	if err != nil {
		return GetMusclesByGroupTypeResp{}, fmt.Errorf("failed to read muscles: %w", err)
	}

	return GetMusclesByGroupTypeResp{Muscle: m}, nil
}

func init() {
	register(&GetMusclesByGroupTypeQuery{})
}
