package commands

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
)

type CreateEquipmentCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Attachment  []*int `json:"attachment"` // Attachment ids
}

type CreateEquipmentResp struct{}

func (cmd *CreateEquipmentCommand) Handle(ctx context.Context) (any, error) {
	equipmentType, err := equipment.NewEquipmentType(cmd.Type)
	if err != nil {
		return nil, fmt.Errorf("failed to create equipment type: %w", err)
	}

	eq, err := equipment.New(cmd.Name, cmd.Description, equipmentType)
	if err != nil {
		return CreateEquipmentResp{}, fmt.Errorf("failed to create equipment: %w", err)
	}

	equipmentID, err := write.Equipment.Add(ctx, eq)
	if err != nil {
		return CreateEquipmentResp{}, fmt.Errorf("failed to insert equipment: %w", err)
	}

	if cmd.Attachment != nil {
		for _, id := range cmd.Attachment {
			_, err := read.Attachment.GetByID(ctx, *id)
			if err != nil {
				return nil, fmt.Errorf("failed to read attachment: %w", err)
			}

			equipmentAttachment := equipment.NewEquipmentAttachment(equipmentID, *id)
			err = write.Equipment.Attachment.Add(ctx, equipmentAttachment)
			if err != nil {
				return nil, fmt.Errorf("failed to insert equipment attachment: %w", err)
			}
		}
	}

	return CreateEquipmentResp{}, nil
}

func init() {
	register(&CreateEquipmentCommand{})
}
