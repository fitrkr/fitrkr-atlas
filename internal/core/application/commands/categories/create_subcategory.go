package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateSubcategoryCommand struct {
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
	Write      ports.SubcategoryWrite
	Read       ports.CategoryRead
}

type CreateSubcategoryResp struct{}

func (cmd *CreateSubcategoryCommand) Handle(ctx context.Context) (any, error) {
	_, err := cmd.Read.GetByID(ctx, cmd.CategoryID)
	if err != nil {
		logr.Get().Errorf("failed to get category : %v", err)
		return CreateSubcategoryResp{}, fmt.Errorf("failed to get category : %w", err)
	}

	sc, err := category.NewSubcategory(cmd.Name, cmd.CategoryID)
	if err != nil {
		logr.Get().Errorf("failed to create new subcategory: %v", err)
		return CreateSubcategoryResp{}, fmt.Errorf("failed to create new subcategory: %w", err)
	}

	err = cmd.Write.Add(ctx, sc)
	if err != nil {
		logr.Get().Errorf("failed to add subcategory: %v", err)
		return CreateSubcategoryResp{}, fmt.Errorf("failed to add subcategory: %w", err)
	}

	return CreateSubcategoryResp{}, nil
}
