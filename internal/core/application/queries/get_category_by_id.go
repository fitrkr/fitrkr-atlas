package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
)

type GetCategoryByIDQuery struct {
	ID int `json:"id"`
}

type GetCategoryByIDResp struct {
	Category *category.Category
}

func (qry *GetCategoryByIDQuery) Handle(ctx context.Context) (any, error) {
	c, err := read.Category.GetByID(ctx, qry.ID)
	if err != nil {
		return GetCategoryByIDResp{}, fmt.Errorf("failed to read category: %w", err)
	}

	return GetCategoryByIDResp{Category: c}, nil
}

func init() {
	register(&GetCategoryByIDQuery{})
}
