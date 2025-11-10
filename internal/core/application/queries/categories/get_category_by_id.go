package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetCategoryByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetCategoryByIDResp struct {
	Category *category.Category
}

func (qry *GetCategoryByIDQuery) Handle(ctx context.Context) (any, error) {
	c, err := qry.Read.Category.GetByID(ctx, qry.ID)
	if err != nil {
		return GetCategoryByIDResp{}, fmt.Errorf("failed to get category: %w", err)
	}

	return GetCategoryByIDResp{Category: c}, nil
}
