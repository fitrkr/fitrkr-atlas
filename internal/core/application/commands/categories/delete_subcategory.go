package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteSubcategoryCommand struct {
	ID    int `json:"id"`
	Write ports.Write
}

type DeleteSubcategoryResp struct{}

func (cmd *DeleteSubcategoryCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Category.Subcategory.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrSubcategoryNotFound {
			return DeleteSubcategoryResp{}, ports.ErrSubcategoryNotFound
		}
		return DeleteSubcategoryResp{}, fmt.Errorf("failed to delete subcategory: %w", err)
	}

	return DeleteSubcategoryResp{}, nil
}
