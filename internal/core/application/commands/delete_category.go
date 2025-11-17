package commands

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteCategoryCommand struct {
	ID int `json:"id"`
}

type DeleteCategoryResp struct{}

func (cmd *DeleteCategoryCommand) Handle(ctx context.Context) (any, error) {
	err := write.Category.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrCategoryNotFound {
			return DeleteCategoryResp{}, ports.ErrCategoryNotFound
		}
		return DeleteCategoryResp{}, fmt.Errorf("failed to delete category: %w", err)
	}

	return DeleteCategoryResp{}, nil
}

func init() {
	register(&DeleteCategoryCommand{})
}
