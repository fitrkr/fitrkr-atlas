package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetMusclesByMuscleGroupIDQuery struct {
	MuscleGroupID int `json:"muscle_group_id"`
	Read          ports.Read
}

type GetMusclesByMuscleGroupIDResp struct {
	Muscles []muscle.Muscle
}

func (qry *GetMusclesByMuscleGroupIDQuery) Handle(ctx context.Context) (any, error) {
	muscles, err := qry.Read.Muscle.GetByMuscleGroupID(ctx, qry.MuscleGroupID)
	if err != nil {
		return GetMusclesByMuscleGroupIDResp{}, fmt.Errorf("failed to get muscles: %w", err)
	}

	return GetMusclesByMuscleGroupIDResp{Muscles: muscles}, nil
}
