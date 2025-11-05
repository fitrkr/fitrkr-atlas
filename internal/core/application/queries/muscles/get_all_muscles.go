package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllMusclesQuery struct {
	Read ports.MuscleRead
}

type GetAllMusclesResp struct {
	Muscles []muscle.Muscle
}

func (qry *GetAllMusclesQuery) Handle(ctx context.Context) (any, error) {
	muscles, err := qry.Read.GetAll(ctx)
	if err != nil {
		logr.Get().Errorf("failed to get muscles: %v", err)
		return GetAllMusclesResp{}, fmt.Errorf("failed to get muscles: %w", err)
	}

	return GetAllMusclesResp{Muscles: muscles}, nil
}
