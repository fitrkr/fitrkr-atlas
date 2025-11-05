package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllCategoriesQuery struct {
	Read ports.CategoryRead
}

type GetAllCategoriesResp struct {
	Categories []category.Category
}

func (qry *GetAllCategoriesQuery) Handle(ctx context.Context) (any, error) {
	categories, err := qry.Read.GetAll(ctx)
	if err != nil {
		logr.Get().Errorf("failed to get categories: %v", err)
		return GetAllCategoriesResp{}, fmt.Errorf("failed to get categories: %w", err)
	}

	return GetAllCategoriesResp{Categories: categories}, nil
}
