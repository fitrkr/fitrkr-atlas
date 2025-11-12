package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetMusclesByGroupTypeQuery struct {
	GroupType string `json:"group_type"`
	Read      ports.Read
}

type GetMusclesByGroupTypeResp struct {
	Muscle []*muscle.Muscle
}

func (qry *GetMusclesByGroupTypeQuery) Handle(ctx context.Context) (any, error) {
	m, err := qry.Read.Muscle.GetByGroupType(ctx, qry.GroupType)
	if err != nil {
		return GetMusclesByGroupTypeResp{}, fmt.Errorf("failed to read muscles: %w", err)
	}

	return GetMusclesByGroupTypeResp{Muscle: m}, nil
}
