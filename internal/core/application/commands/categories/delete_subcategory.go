package categories

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteSubcategoryCommand struct {
	ID    int `json:"id"`
	Write ports.SubcategoryWrite
}

type DeleteSubcategoryResp struct{}

func (cmd *DeleteSubcategoryCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrSubcategoryNotFound {
			logr.Get().Error("subcategory not found")
			return DeleteSubcategoryResp{}, ports.ErrSubcategoryNotFound
		}
		logr.Get().Errorf("failed to delete subcategory: %v", err)
		return DeleteSubcategoryResp{}, fmt.Errorf("failed to delete subcategory: %w", err)
	}

	logr.Get().Info("subcategory deleted successfully")
	return DeleteSubcategoryResp{}, nil
}
