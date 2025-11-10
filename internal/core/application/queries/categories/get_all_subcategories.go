package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllSubcategoriesQuery struct {
	Read ports.Read
}

type GetAllSubcategoriesResp struct {
	Subcategories []category.Subcategory
}

func (qry *GetAllSubcategoriesQuery) Handle(ctx context.Context) (any, error) {
	subcategories, err := qry.Read.Category.Subcategory.GetAll(ctx)
	if err != nil {
		return GetAllSubcategoriesResp{}, fmt.Errorf("failed to get subcategories: %w", err)
	}

	return GetAllSubcategoriesResp{Subcategories: subcategories}, nil
}
