package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteCategoryCommand struct {
	ID    int `json:"id"`
	Write ports.CategoryWrite
}

type DeleteCategoryResp struct{}

func (cmd *DeleteCategoryCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrCategoryNotFound {
			logr.Get().Error("category not found")
			return DeleteCategoryResp{}, ports.ErrCategoryNotFound
		}
		logr.Get().Errorf("failed to delete category: %v", err)
		return DeleteCategoryResp{}, fmt.Errorf("failed to delete category: %w", err)
	}

	logr.Get().Info("category deleted successfully")
	return DeleteCategoryResp{}, nil
}
