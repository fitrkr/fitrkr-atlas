package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/fitrkr/atlas/internal/core/domain/category"
)

type UpdateCategoryCommand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type UpdateCategoryResp struct{}

func (cmd *UpdateCategoryCommand) Handle(ctx context.Context) (any, error) {
	existing, err := read.Category.GetByID(ctx, cmd.ID)
	if err != nil {
		return UpdateCategoryResp{}, fmt.Errorf("failed to read category: %w", err)
	}

	if cmd.Name != "" {
		name := strings.TrimSpace(strings.ToLower(cmd.Name))
		existing.Name = name
	}

	if cmd.Type != "" && cmd.Type != existing.Type {
		categoryType, err := category.NewCategoryType(cmd.Type)
		if err != nil {
			return CreateCategoryResp{}, err
		}

		existing.Type = categoryType.ToString()
	}

	existing.Touch()

	err = write.Category.Update(ctx, *existing)
	if err != nil {
		return UpdateCategoryResp{}, fmt.Errorf("failed to update category: %w", err)
	}

	return UpdateCategoryResp{}, nil
}

func init() {
	register(&UpdateCategoryCommand{})
}
