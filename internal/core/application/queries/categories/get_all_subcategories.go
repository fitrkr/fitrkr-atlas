package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllSubcategoriesQuery struct {
	Read ports.SubcategoryRead
}

type GetAllSubcategoriesResp struct {
	Subcategories []category.Subcategory
}

func (qry *GetAllSubcategoriesQuery) Handle(ctx context.Context) (any, error) {
	subcategories, err := qry.Read.GetAll(ctx)
	if err != nil {
		logr.Get().Errorf("failed to get subcategories: %v", err)
		return GetAllSubcategoriesResp{}, fmt.Errorf("failed to get subcategories: %w", err)
	}

	return GetAllSubcategoriesResp{Subcategories: subcategories}, nil
}
