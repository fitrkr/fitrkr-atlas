package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateSubcategoryCommand struct {
	ID         int    `json:"id"`
	CategoryID *int   `json:"category_id"`
	Name       string `json:"name"`
	Write      ports.Write
	Read       ports.Read
}

type UpdateSubcategoryResp struct{}

func (cmd *UpdateSubcategoryCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.Category.Subcategory.GetByID(ctx, cmd.ID)
	if err != nil {
		return UpdateSubcategoryResp{}, fmt.Errorf("failed to get subcategory: %w", err)
	}

	if cmd.Name != "" {
		existing.Name = cmd.Name
	}
	if cmd.CategoryID != nil {
		_, err := cmd.Read.Category.GetByID(ctx, *cmd.CategoryID)
		if err != nil {
			return UpdateSubcategoryResp{}, fmt.Errorf("failed to get category: %w", err)
		}
		existing.CategoryID = *cmd.CategoryID
	}
	existing.Touch()

	err = cmd.Write.Category.Subcategory.Update(ctx, *existing)
	if err != nil {
		return UpdateSubcategoryResp{}, fmt.Errorf("failed to update subcategory: %w", err)
	}

	return UpdateSubcategoryResp{}, nil
}
