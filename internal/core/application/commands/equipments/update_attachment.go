package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateAttachmentCommand struct {
	ID            int    `json:"id"`
	EquipmentID   int    `json:"equipment_id"`
	Name          string `json:"name"`
	Write         ports.EquipmentAttachmentWrite
	Read          ports.EquipmentAttachmentRead
	ReadEquipment ports.EquipmentRead
}

type UpdateAttachmentResp struct{}

func (cmd *UpdateAttachmentCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.GetByID(ctx, cmd.ID)
	if err != nil {
		logr.Get().Errorf("failed to get attachment: %v", err)
		return UpdateAttachmentResp{}, fmt.Errorf("failed to get attachment: %w", err)
	}

	if cmd.Name != "" {
		existing.Name = cmd.Name
	}
	if cmd.EquipmentID > 0 {
		_, err := cmd.ReadEquipment.GetByID(ctx, cmd.EquipmentID)
		if err != nil {
			logr.Get().Errorf("failed to validate equipment: %v", err)
			return UpdateAttachmentResp{}, fmt.Errorf("failed to validate equipment: %w", err)
		}
		existing.EquipmentID = cmd.EquipmentID
	}

	existing.Touch()

	err = cmd.Write.Update(ctx, *existing)
	if err != nil {
		logr.Get().Errorf("failed to update attachment: %v", err)
		return UpdateAttachmentResp{}, fmt.Errorf("failed to update attachment: %w", err)
	}

	return UpdateAttachmentResp{}, nil
}
