package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetCategoryByIDQuery struct {
	ID   int `json:"id"`
	Read ports.CategoryRead
}

type GetCategoryByIDResp struct {
	Category *category.Category
}

func (qry *GetCategoryByIDQuery) Handle(ctx context.Context) (any, error) {
	c, err := qry.Read.GetByID(ctx, qry.ID)
	if err != nil {
		logr.Get().Errorf("failed to get category: %v", err)
		return GetCategoryByIDResp{}, fmt.Errorf("failed to get category: %w", err)
	}

	return GetCategoryByIDResp{Category: c}, nil
}
