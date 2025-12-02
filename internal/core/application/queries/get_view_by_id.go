package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/view"
)

type GetViewByIDQuery struct {
	ID int `json:"id"`
}

type GetViewByIDResp struct {
	View *view.View
}

func (qry *GetViewByIDQuery) Handle(ctx context.Context) (any, error) {
	v, err := read.View.GetByID(ctx, qry.ID)
	if err != nil {
		return GetViewByIDResp{}, fmt.Errorf("failed to get view: %w", err)
	}

	return GetViewByIDResp{View: v}, nil
}

func init() {
	register(&GetViewByIDQuery{})
}
