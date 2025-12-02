package commands

import (
	"context"
	"fmt"

	"github.com/fitrkr/atlas/internal/core/domain/equipment"
)

type CreateEquipmentAttachmentCommand struct {
	EquipmentID  int `json:"equipment_id"`
	AttachmentID int `json:"attachment_id"`
}

type CreateEquipmentAttachmentResp struct{}

func (cmd *CreateEquipmentAttachmentCommand) Handle(ctx context.Context) (any, error) {
	_, err := read.Attachment.GetByID(ctx, cmd.AttachmentID)
	if err != nil {
		return CreateEquipmentAttachmentResp{}, fmt.Errorf("failed to read attachment: %w", err)
	}
	_, err = read.Equipment.GetByID(ctx, cmd.EquipmentID)
	if err != nil {
		return CreateEquipmentAttachmentResp{}, fmt.Errorf("failed to read equipment: %w", err)
	}

	equipmentAttachment := equipment.NewEquipmentAttachment(cmd.EquipmentID, cmd.AttachmentID)

	err = write.Equipment.Attachment.Add(ctx, equipmentAttachment)
	if err != nil {
		return CreateEquipmentAttachmentResp{}, fmt.Errorf("failed to insert equipment attachment: %w", err)
	}

	return CreateEquipmentAttachmentResp{}, nil
}

func init() {
	register(&CreateEquipmentAttachmentCommand{})
}
