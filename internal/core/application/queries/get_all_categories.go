package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/category"
)

type GetAllCategoriesQuery struct{}

type GetAllCategoriesResp struct {
	Categories []*category.Category
}

func (qry *GetAllCategoriesQuery) Handle(ctx context.Context) (any, error) {
	categories, err := read.Category.GetAll(ctx)
	if err != nil {
		return GetAllCategoriesResp{}, fmt.Errorf("failed to read categories: %w", err)
	}

	return GetAllCategoriesResp{Categories: categories}, nil
}

func init() {
	register(&GetAllCategoriesQuery{})
}
