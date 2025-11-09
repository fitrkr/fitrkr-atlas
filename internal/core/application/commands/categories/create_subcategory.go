package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateSubcategoryCommand struct {
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
	Write      ports.Write
	Read       ports.Read
}

type CreateSubcategoryResp struct{}

func (cmd *CreateSubcategoryCommand) Handle(ctx context.Context) (any, error) {
	_, err := cmd.Read.Category.GetByID(ctx, cmd.CategoryID)
	if err != nil {
		return CreateSubcategoryResp{}, fmt.Errorf("failed to get category : %w", err)
	}

	subcategory, err := category.NewSubcategory(cmd.Name, cmd.CategoryID)
	if err != nil {
		return CreateSubcategoryResp{}, fmt.Errorf("failed to create new subcategory: %w", err)
	}

	err = cmd.Write.Category.Subcategory.Add(ctx, subcategory)
	if err != nil {
		return CreateSubcategoryResp{}, fmt.Errorf("failed to add subcategory: %w", err)
	}

	return CreateSubcategoryResp{}, nil
}
