package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetMuscleGroupByIDQuery struct {
	ID   int `json:"id"`
	Read ports.MuscleGroupRead
}

type GetMuscleGroupByIDResp struct {
	MuscleGroup *muscle.Group
}

func (qry *GetMuscleGroupByIDQuery) Handle(ctx context.Context) (any, error) {
	mg, err := qry.Read.GetByID(ctx, qry.ID)
	if err != nil {
		logr.Get().Errorf("failed to get muscle group: %v", err)
		return GetMuscleGroupByIDResp{}, fmt.Errorf("failed to get muscle group: %w", err)
	}

	return GetMuscleGroupByIDResp{MuscleGroup: mg}, nil
}
