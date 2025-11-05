package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteAttachmentCommand struct {
	ID    int `json:"id"`
	Write ports.EquipmentAttachmentWrite
}

type DeleteAttachmentResp struct{}

func (cmd *DeleteAttachmentCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrAttachmentNotFound {
			logr.Get().Error("attachment not found")
			return DeleteAttachmentResp{}, ports.ErrAttachmentNotFound
		}
		logr.Get().Errorf("failed to delete attachment: %v", err)
		return DeleteAttachmentResp{}, fmt.Errorf("failed to delete attachment: %w", err)
	}

	logr.Get().Info("Attachment deleted successfully")

	return DeleteAttachmentResp{}, nil
}
