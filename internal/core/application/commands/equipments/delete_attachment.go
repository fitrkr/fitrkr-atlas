package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteAttachmentCommand struct {
	ID    int `json:"id"`
	Write ports.Write
}

type DeleteAttachmentResp struct{}

func (cmd *DeleteAttachmentCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Attachment.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrAttachmentNotFound {
			return DeleteAttachmentResp{}, ports.ErrAttachmentNotFound
		}
		return DeleteAttachmentResp{}, fmt.Errorf("failed to delete attachment: %w", err)
	}

	return DeleteAttachmentResp{}, nil
}
