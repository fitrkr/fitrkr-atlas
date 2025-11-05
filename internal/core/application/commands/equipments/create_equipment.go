// Package equipments
package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateEquipmentCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Write       ports.EquipmentWrite
	Read        ports.EquipmentRead
}

type CreateEquipmentResp struct{}

func (cmd *CreateEquipmentCommand) Handle(ctx context.Context) (any, error) {
	eq, err := equipment.New(cmd.Name, cmd.Description)
	if err != nil {
		logr.Get().Errorf("failed to create new equipment: %v", err)
		return CreateEquipmentResp{}, fmt.Errorf("failed to create new equipment: %w", err)
	}

	err = cmd.Write.Add(ctx, eq)
	if err != nil {
		logr.Get().Errorf("failed to add equipment to db: %v", err)
		return CreateEquipmentResp{}, fmt.Errorf("failed to add equipment to db: %w", err)
	}

	return CreateEquipmentResp{}, nil
}
