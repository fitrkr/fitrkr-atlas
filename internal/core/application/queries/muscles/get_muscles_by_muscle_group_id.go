package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetMusclesByMuscleGroupIDQuery struct {
	MuscleGroupID int `json:"muscle_group_id"`
	Read          ports.MuscleRead
}

type GetMusclesByMuscleGroupIDResp struct {
	Muscles []muscle.Muscle
}

func (qry *GetMusclesByMuscleGroupIDQuery) Handle(ctx context.Context) (any, error) {
	muscles, err := qry.Read.GetByMuscleGroupID(ctx, qry.MuscleGroupID)
	if err != nil {
		logr.Get().Errorf("failed to get muscles: %v", err)
		return GetMusclesByMuscleGroupIDResp{}, fmt.Errorf("failed to get muscles: %w", err)
	}

	return GetMusclesByMuscleGroupIDResp{Muscles: muscles}, nil
}
