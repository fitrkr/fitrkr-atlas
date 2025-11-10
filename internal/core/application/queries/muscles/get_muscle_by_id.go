package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetMuscleByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetMuscleByIDResp struct {
	Muscle *muscle.Muscle
}

func (qry *GetMuscleByIDQuery) Handle(ctx context.Context) (any, error) {
	m, err := qry.Read.Muscle.GetByID(ctx, qry.ID)
	if err != nil {
		return GetMuscleByIDResp{}, fmt.Errorf("failed to get muscle: %w", err)
	}

	return GetMuscleByIDResp{Muscle: m}, nil
}
