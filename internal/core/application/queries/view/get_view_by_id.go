package view

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/view"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetViewByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetViewByIDResp struct {
	View *view.View
}

func (qry *GetViewByIDQuery) Handle(ctx context.Context) (any, error) {
	v, err := qry.Read.View.GetByID(ctx, qry.ID)
	if err != nil {
		return GetViewByIDResp{}, fmt.Errorf("failed to get view: %w", err)
	}

	return GetViewByIDResp{View: v}, nil
}
