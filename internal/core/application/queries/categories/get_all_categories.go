// Package categories
package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllCategoriesQuery struct {
	Read ports.Read
}

type GetAllCategoriesResp struct {
	Categories []category.Category
}

func (qry *GetAllCategoriesQuery) Handle(ctx context.Context) (any, error) {
	categories, err := qry.Read.Category.GetAll(ctx)
	if err != nil {
		return GetAllCategoriesResp{}, fmt.Errorf("failed to get categories: %w", err)
	}

	return GetAllCategoriesResp{Categories: categories}, nil
}
