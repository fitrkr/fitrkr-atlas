package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateSubcategoryCommand struct {
	ID           int    `json:"id"`
	CategoryID   int    `json:"category_id"`
	Name         string `json:"name"`
	Write        ports.SubcategoryWrite
	Read         ports.SubcategoryRead
	ReadCategory ports.CategoryRead
}

type UpdateSubcategoryResp struct{}

func (cmd *UpdateSubcategoryCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.GetByID(ctx, cmd.ID)
	if err != nil {
		logr.Get().Errorf("failed to get subcategory: %v", err)
		return UpdateSubcategoryResp{}, fmt.Errorf("failed to get subcategory: %w", err)
	}

	if cmd.Name != "" {
		existing.Name = cmd.Name
	}
	if cmd.CategoryID > 0 {
		_, err := cmd.ReadCategory.GetByID(ctx, cmd.CategoryID)
		if err != nil {
			logr.Get().Errorf("failed to get category: %v", err)
			return UpdateSubcategoryResp{}, fmt.Errorf("failed to get category: %w", err)
		}
		existing.CategoryID = cmd.CategoryID
	}
	existing.Touch()

	err = cmd.Write.Update(ctx, *existing)
	if err != nil {
		logr.Get().Errorf("failed to update subcategory: %v", err)
		return UpdateSubcategoryResp{}, fmt.Errorf("failed to update subcategory: %w", err)
	}

	return UpdateSubcategoryResp{}, nil
}
