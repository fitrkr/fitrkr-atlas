package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllMuscleGroupsQuery struct {
	Read ports.MuscleGroupRead
}

type GetAllMuscleGroupsResp struct {
	MuscleGroups []muscle.Group
}

func (qry *GetAllMuscleGroupsQuery) Handle(ctx context.Context) (any, error) {
	muscleGroups, err := qry.Read.GetAll(ctx)
	if err != nil {
		logr.Get().Errorf("failed to get muscle groups: %v", err)
		return GetAllMuscleGroupsResp{}, fmt.Errorf("failed to get muscle groups: %w", err)
	}

	return GetAllMuscleGroupsResp{MuscleGroups: muscleGroups}, nil
}
