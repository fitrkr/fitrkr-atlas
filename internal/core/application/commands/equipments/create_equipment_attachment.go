package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateEquipmentAttachmentCommand struct {
	EquipmentID  int `json:"equipment_id"`
	AttachmentID int `json:"attachment_id"`
	Write        ports.Write
	Read         ports.Read
}

type CreateEquipmentAttachmentResp struct{}

func (cmd *CreateEquipmentAttachmentCommand) Handle(ctx context.Context) (any, error) {
	_, err := cmd.Read.Attachment.GetByID(ctx, cmd.AttachmentID)
	if err != nil {
		return CreateEquipmentAttachmentResp{}, fmt.Errorf("failed to validate attachment: %w", err)
	}
	_, err = cmd.Read.Exercise.GetByID(ctx, cmd.EquipmentID)
	if err != nil {
		return CreateEquipmentAttachmentResp{}, fmt.Errorf("failed to validate equipment: %w", err)
	}

	equipmentAttachment := equipment.NewEquipmentAttachment(cmd.EquipmentID, cmd.AttachmentID)

	err = cmd.Write.Equipment.Attachment.Add(ctx, equipmentAttachment)
	if err != nil {
		return CreateEquipmentAttachmentResp{}, fmt.Errorf("failed to add equipment attachment: %w", err)
	}

	return CreateEquipmentAttachmentResp{}, nil
}
