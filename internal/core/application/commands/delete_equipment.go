package commands

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type DeleteEquipmentCommand struct {
	ID int `json:"id"`
}

type DeleteEquipmentResp struct{}

func (cmd *DeleteEquipmentCommand) Handle(ctx context.Context) (any, error) {
	err := write.Equipment.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrEquipmentNotFound {
			return DeleteEquipmentResp{}, ports.ErrEquipmentNotFound
		}
		return DeleteEquipmentResp{}, fmt.Errorf("failed to delete equipment: %w", err)
	}

	return DeleteEquipmentResp{}, nil
}

func init() {
	register(&DeleteEquipmentCommand{})
}
