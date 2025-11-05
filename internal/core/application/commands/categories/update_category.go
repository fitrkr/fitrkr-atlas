package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateCategoryCommand struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Write ports.CategoryWrite
	Read  ports.CategoryRead
}

type UpdateCategoryResp struct{}

func (cmd *UpdateCategoryCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.GetByID(ctx, cmd.ID)
	if err != nil {
		logr.Get().Errorf("failed to get category: %v", err)
		return UpdateCategoryResp{}, fmt.Errorf("failed to get category: %w", err)
	}

	if cmd.Name != "" {
		name, err := category.NewCategoryType(cmd.Name)
		if err != nil {
			logr.Get().Errorf("failed to create new category type: %v", err)
			return UpdateCategoryResp{}, fmt.Errorf("failed to create new category type: %w", err)
		}
		existing.Name = name
	}

	existing.Touch()

	err = cmd.Write.Update(ctx, *existing)
	if err != nil {
		logr.Get().Errorf("failed to update category: %v", err)
		return UpdateCategoryResp{}, fmt.Errorf("failed to update category: %w", err)
	}

	return UpdateCategoryResp{}, nil
}
