package muscles

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllMusclesQuery struct {
	Read ports.Read
}

type GetAllMusclesResp struct {
	Muscles []muscle.Muscle
}

func (qry *GetAllMusclesQuery) Handle(ctx context.Context) (any, error) {
	muscles, err := qry.Read.Muscle.GetAll(ctx)
	if err != nil {
		return GetAllMusclesResp{}, fmt.Errorf("failed to get muscles: %w", err)
	}

	return GetAllMusclesResp{Muscles: muscles}, nil
}
