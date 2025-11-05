package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetSubcategoryByIDQuery struct {
	ID   int `json:"id"`
	Read ports.SubcategoryRead
}

type GetSubcategoryByIDResp struct {
	Subcategory *category.Subcategory
}

func (qry *GetSubcategoryByIDQuery) Handle(ctx context.Context) (any, error) {
	subcategory, err := qry.Read.GetByID(ctx, qry.ID)
	if err != nil {
		logr.Get().Errorf("failed to get subcategory: %v", err)
		return GetSubcategoryByIDResp{}, fmt.Errorf("failed to get subcategory: %w", err)
	}

	return GetSubcategoryByIDResp{Subcategory: subcategory}, nil
}
