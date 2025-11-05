package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetMuscleByIDQuery struct {
	ID   int `json:"id"`
	Read ports.MuscleRead
}

type GetMuscleByIDResp struct {
	Muscle *muscle.Muscle
}

func (qry *GetMuscleByIDQuery) Handle(ctx context.Context) (any, error) {
	m, err := qry.Read.GetByID(ctx, qry.ID)
	if err != nil {
		logr.Get().Errorf("failed to get muscle: %v", err)
		return GetMuscleByIDResp{}, fmt.Errorf("failed to get muscle: %w", err)
	}

	return GetMuscleByIDResp{Muscle: m}, nil
}
