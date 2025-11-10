package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAliasByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetAliasByIDResp struct {
	Alias *exercise.Alias
}

func (qry *GetAliasByIDQuery) Handle(ctx context.Context) (any, error) {
	alias, err := qry.Read.Exercise.Alias.GetByID(ctx, qry.ID)
	if err != nil {
		return GetAliasByIDResp{}, fmt.Errorf("failed to get exercise alias: %w", err)
	}

	return GetAliasByIDResp{Alias: alias}, nil
}
