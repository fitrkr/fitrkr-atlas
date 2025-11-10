package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetSubcategoryByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetSubcategoryByIDResp struct {
	Subcategory *category.Subcategory
}

func (qry *GetSubcategoryByIDQuery) Handle(ctx context.Context) (any, error) {
	subcategory, err := qry.Read.Category.Subcategory.GetByID(ctx, qry.ID)
	if err != nil {
		return GetSubcategoryByIDResp{}, fmt.Errorf("failed to get subcategory: %w", err)
	}

	return GetSubcategoryByIDResp{Subcategory: subcategory}, nil
}
