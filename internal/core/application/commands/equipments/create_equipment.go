// Package equipments
package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type EquipmentAttachmentReq struct {
	AttachmentID int
}

type CreateEquipmentCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Attachment  []*EquipmentAttachmentReq
	Write       ports.Write
	Read        ports.Read
}

type CreateEquipmentResp struct{}

func (cmd *CreateEquipmentCommand) Handle(ctx context.Context) (any, error) {
	equipmentType, err := equipment.NewEquipmentType(cmd.Type)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new equipment type: %w", err)
	}

	eq, err := equipment.New(cmd.Name, cmd.Description, equipmentType)
	if err != nil {
		return CreateEquipmentResp{}, fmt.Errorf("failed to create new equipment: %w", err)
	}

	id, err := cmd.Write.Equipment.Add(ctx, eq)
	if err != nil {
		return CreateEquipmentResp{}, fmt.Errorf("failed to add equipment to db: %w", err)
	}

	if cmd.Attachment != nil {
		for _, a := range cmd.Attachment {
			equipmentAttachment := equipment.NewEquipmentAttachment(id, a.AttachmentID)
			cmd.Write.Equipment.Attachment.Add(ctx, equipmentAttachment)
		}
	}

	return CreateEquipmentResp{}, nil
}
