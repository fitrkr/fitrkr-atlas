// Package categories
package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateCategoryCommand struct {
	Name  string `json:"name"`
	Write ports.Write
}

type CreateCategoryResp struct{}

func (cmd *CreateCategoryCommand) Handle(ctx context.Context) (any, error) {
	name, err := category.NewCategoryType(cmd.Name)
	if err != nil {
		return CreateCategoryResp{}, fmt.Errorf("failed to create new category type: %w", err)
	}

	c, err := category.New(name)
	if err != nil {
		return CreateCategoryResp{}, fmt.Errorf("failed to create new category: %w", err)
	}

	err = cmd.Write.Category.Add(ctx, c)
	if err != nil {
		return CreateCategoryResp{}, fmt.Errorf("failed to add category: %w", err)
	}

	return CreateCategoryResp{}, nil
}
