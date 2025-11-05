package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetSubcategoriesByCategoryIDQuery struct {
	CategoryID int `json:"category_id"`
	Read       ports.SubcategoryRead
}

type GetSubcategoriesByCategoryIDResp struct {
	Subcategories []category.Subcategory
}

func (qry *GetSubcategoriesByCategoryIDQuery) Handle(ctx context.Context) (any, error) {
	subcategories, err := qry.Read.GetByCategoryID(ctx, qry.CategoryID)
	if err != nil {
		logr.Get().Errorf("failed to get subcategories: %v", err)
		return GetSubcategoriesByCategoryIDResp{}, fmt.Errorf("failed to get subcategories: %w", err)
	}

	return GetSubcategoriesByCategoryIDResp{Subcategories: subcategories}, nil
}
