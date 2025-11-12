package categories

import (
	"context"
	"fmt"
	"strings"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateCategoryCommand struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Write ports.Write
	Read  ports.Read
}

type UpdateCategoryResp struct{}

func (cmd *UpdateCategoryCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.Category.GetByID(ctx, cmd.ID)
	if err != nil {
		return UpdateCategoryResp{}, fmt.Errorf("failed to read category: %w", err)
	}

	if cmd.Name != "" {
		name := strings.TrimSpace(strings.ToLower(cmd.Name))
		existing.Name = name
	}

	if cmd.Type != "" && cmd.Type != existing.Type.ToString() {
		categoryType, err := category.NewCategoryType(cmd.Type)
		if err != nil {
			return CreateCategoryResp{}, err
		}

		existing.Type = categoryType
	}

	existing.Touch()

	err = cmd.Write.Category.Update(ctx, *existing)
	if err != nil {
		return UpdateCategoryResp{}, fmt.Errorf("failed to update category: %w", err)
	}

	return UpdateCategoryResp{}, nil
}
