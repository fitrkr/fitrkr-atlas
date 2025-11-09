package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateCategoryCommand struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Write ports.Write
	Read  ports.Read
}

type UpdateCategoryResp struct{}

func (cmd *UpdateCategoryCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.Category.GetByID(ctx, cmd.ID)
	if err != nil {
		return UpdateCategoryResp{}, fmt.Errorf("failed to get category: %w", err)
	}

	if cmd.Name != "" {
		name, err := category.NewCategoryType(cmd.Name)
		if err != nil {
			return UpdateCategoryResp{}, fmt.Errorf("failed to create new category type: %w", err)
		}
		existing.Name = name
	}

	existing.Touch()

	err = cmd.Write.Category.Update(ctx, *existing)
	if err != nil {
		return UpdateCategoryResp{}, fmt.Errorf("failed to update category: %w", err)
	}

	return UpdateCategoryResp{}, nil
}
