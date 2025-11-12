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
	Type  string `json:"type"`
	Write ports.Write
}

type CreateCategoryResp struct{}

func (cmd *CreateCategoryCommand) Handle(ctx context.Context) (any, error) {
	categoryType, err := category.NewCategoryType(cmd.Type)
	if err != nil {
		return CreateCategoryResp{}, err
	}

	c, err := category.New(cmd.Name, categoryType)
	if err != nil {
		return CreateCategoryResp{}, fmt.Errorf("failed to create new category: %w", err)
	}

	err = cmd.Write.Category.Add(ctx, c)
	if err != nil {
		return CreateCategoryResp{}, fmt.Errorf("failed to insert category: %w", err)
	}

	return CreateCategoryResp{}, nil
}
