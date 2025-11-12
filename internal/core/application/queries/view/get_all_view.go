// Package view
package view

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/view"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllViewQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetAllViewResp struct {
	View []*view.View
}

func (qry *GetAllViewQuery) Handle(ctx context.Context) (any, error) {
	v, err := qry.Read.View.GetAll(ctx)
	if err != nil {
		return GetAllViewResp{}, fmt.Errorf("failed to get view: %w", err)
	}

	return GetAllViewResp{View: v}, nil
}
