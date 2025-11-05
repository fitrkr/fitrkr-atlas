package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateEquipmentCommand struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Write       ports.EquipmentWrite
	Read        ports.EquipmentRead
}

type UpdateEquipmentResp struct{}

func (cmd *UpdateEquipmentCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.GetByID(ctx, cmd.ID)
	if err != nil {
		logr.Get().Errorf("failed to get equipment: %v", err)
		return UpdateEquipmentResp{}, fmt.Errorf("failed to get equipment: %w", err)
	}

	if cmd.Name != "" {
		existing.Name = cmd.Name
	}
	if cmd.Description != "" {
		existing.Description = &cmd.Description
	}
	existing.Touch()

	err = cmd.Write.Update(ctx, *existing)
	if err != nil {
		logr.Get().Errorf("failed to update equipment: %v", err)
		return UpdateEquipmentResp{}, fmt.Errorf("failed to update equipment: %w", err)
	}

	return UpdateEquipmentResp{}, nil
}
