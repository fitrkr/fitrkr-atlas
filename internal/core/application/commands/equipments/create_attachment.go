package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateAttachmentCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	EquipmentID int    `json:"equipment_id"`
	Write       ports.EquipmentAttachmentWrite
	Read        ports.EquipmentRead
}

type CreateAttachmentResp struct{}

func (cmd *CreateAttachmentCommand) Handle(ctx context.Context) (any, error) {
	// Check if equipment exists
	_, err := cmd.Read.GetByID(ctx, cmd.EquipmentID)
	if err != nil {
		logr.Get().Errorf("failed to validate equipment: %v", err)
		return CreateAttachmentResp{}, fmt.Errorf("failed to validate equipment: %w", err)
	}

	at, err := equipment.NewAttachment(cmd.Name, cmd.EquipmentID)
	if err != nil {
		logr.Get().Errorf("failed to create new equipment attachment: %v", err)
		return CreateAttachmentResp{}, fmt.Errorf("failed to create new equipment attachment: %w", err)
	}

	err = cmd.Write.Add(ctx, at)
	if err != nil {
		logr.Get().Errorf("failed to add equipment attachment to db: %v", err)
		return CreateAttachmentResp{}, fmt.Errorf("failed to add equipment attachment to db: %w", err)
	}

	return CreateAttachmentResp{}, nil
}
