package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
)

type GetCategoriesByTypeQuery struct {
	Type string `json:"type"`
}

type GetCategoriesByTypeResp struct {
	Category []*category.Category
}

func (qry *GetCategoriesByTypeQuery) Handle(ctx context.Context) (any, error) {
	c, err := read.Category.GetByType(ctx, qry.Type)
	if err != nil {
		return GetCategoriesByTypeResp{}, fmt.Errorf("failed to read category: %w", err)
	}

	return GetCategoriesByTypeResp{Category: c}, nil
}

func init() {
	register(&GetCategoriesByTypeQuery{})
}
