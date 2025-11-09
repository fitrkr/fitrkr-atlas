package equipments

import (
	"context"
	"fmt"
	"strings"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateEquipmentCommand struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Write       ports.Write
	Read        ports.Read
}

type UpdateEquipmentResp struct{}

func (cmd *UpdateEquipmentCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.Equipment.GetByID(ctx, cmd.ID)
	if err != nil {
		return UpdateEquipmentResp{}, fmt.Errorf("failed to get equipment: %w", err)
	}

	if cmd.Name != "" {
		name := strings.TrimSpace(strings.ToLower(cmd.Name))
		existing.Name = name
	}
	if cmd.Description != "" {
		existing.Description = &cmd.Description
	}
	if cmd.Type != "" {
		equipmentType, err := equipment.NewEquipmentType(cmd.Type)
		if err != nil {
			return nil, fmt.Errorf("failed to create a new equipment type: %w", err)
		}
		existing.Type = equipmentType
	}

	existing.Touch()

	err = cmd.Write.Equipment.Update(ctx, *existing)
	if err != nil {
		return UpdateEquipmentResp{}, fmt.Errorf("failed to update equipment: %w", err)
	}

	return UpdateEquipmentResp{}, nil
}
