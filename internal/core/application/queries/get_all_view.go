package queries

import (
	"context"
	"fmt"

	"github.com/fitrkr/atlas/internal/core/domain/view"
)

type GetAllViewQuery struct {
	ID int `json:"id"`
}

type GetAllViewResp struct {
	View []*view.View
}

func (qry *GetAllViewQuery) Handle(ctx context.Context) (any, error) {
	v, err := read.View.GetAll(ctx)
	if err != nil {
		return GetAllViewResp{}, fmt.Errorf("failed to get view: %w", err)
	}

	return GetAllViewResp{View: v}, nil
}

func init() {
	register(&GetAllViewQuery{})
}
