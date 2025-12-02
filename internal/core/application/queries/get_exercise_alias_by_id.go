package queries

import (
	"context"
	"fmt"

	"github.com/fitrkr/atlas/internal/core/domain/exercise"
)

type GetAliasByIDQuery struct {
	ID int `json:"id"`
}

type GetAliasByIDResp struct {
	Alias *exercise.Alias
}

func (qry *GetAliasByIDQuery) Handle(ctx context.Context) (any, error) {
	alias, err := read.Exercise.Alias.GetByID(ctx, qry.ID)
	if err != nil {
		return GetAliasByIDResp{}, fmt.Errorf("failed to get exercise alias: %w", err)
	}

	return GetAliasByIDResp{Alias: alias}, nil
}

func init() {
	register(&GetAliasByIDQuery{})
}
