package commands

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteEquipmentAttachmentCommand struct {
	ID int `json:"id"`
}

type DeleteEquipmentAttachmentResp struct{}

func (cmd *DeleteEquipmentAttachmentCommand) Handle(ctx context.Context) (any, error) {
	err := write.Equipment.Attachment.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrEquipmentAttachmentNotFound {
			return DeleteEquipmentAttachmentResp{}, ports.ErrEquipmentAttachmentNotFound
		}
		return DeleteEquipmentAttachmentResp{}, fmt.Errorf("failed to delete equipment attachment: %w", err)
	}

	return DeleteEquipmentAttachmentResp{}, nil
}

func init() {
	register(&DeleteEquipmentAttachmentCommand{})
}
