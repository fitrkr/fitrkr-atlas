// Package muscles
package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllMuscleGroupsQuery struct {
	Read ports.Read
}

type GetAllMuscleGroupsResp struct {
	MuscleGroups []muscle.Group
}

func (qry *GetAllMuscleGroupsQuery) Handle(ctx context.Context) (any, error) {
	muscleGroups, err := qry.Read.Muscle.Group.GetAll(ctx)
	if err != nil {
		return GetAllMuscleGroupsResp{}, fmt.Errorf("failed to get muscle groups: %w", err)
	}

	return GetAllMuscleGroupsResp{MuscleGroups: muscleGroups}, nil
}
