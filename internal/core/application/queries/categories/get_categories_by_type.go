package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetCategoriesByTypeQuery struct {
	Type string `json:"type"`
	Read ports.Read
}

type GetCategoriesByTypeResp struct {
	Category []*category.Category
}

func (qry *GetCategoriesByTypeQuery) Handle(ctx context.Context) (any, error) {
	c, err := qry.Read.Category.GetByType(ctx, qry.Type)
	if err != nil {
		return GetCategoriesByTypeResp{}, fmt.Errorf("failed to read category: %w", err)
	}

	return GetCategoriesByTypeResp{Category: c}, nil
}
