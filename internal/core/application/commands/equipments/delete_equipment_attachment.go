package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteEquipmentAttachmentCommand struct {
	ID    int `json:"id"`
	Write ports.Write
}

type DeleteEquipmentAttachmentResp struct{}

func (cmd *DeleteEquipmentAttachmentCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Equipment.Attachment.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrEquipmentAttachmentNotFound {
			return DeleteEquipmentAttachmentResp{}, ports.ErrEquipmentAttachmentNotFound
		}
		return DeleteEquipmentAttachmentResp{}, fmt.Errorf("failed to delete equipment attachment: %w", err)
	}

	return DeleteEquipmentAttachmentResp{}, nil
}
