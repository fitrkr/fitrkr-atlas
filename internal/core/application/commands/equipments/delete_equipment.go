package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type DeleteEquipmentCommand struct {
	ID    int `json:"id"`
	Write ports.EquipmentWrite
}

type DeleteEquipmentResp struct{}

func (cmd *DeleteEquipmentCommand) Handle(ctx context.Context) (any, error) {
	err := cmd.Write.Delete(ctx, cmd.ID)
	if err != nil {
		if err == ports.ErrEquipmentNotFound {
			logr.Get().Error("equipment not found")
			return DeleteEquipmentResp{}, ports.ErrEquipmentNotFound
		}
		logr.Get().Errorf("failed to delete equipment: %v", err)
		return DeleteEquipmentResp{}, fmt.Errorf("failed to delete equipment: %w", err)
	}

	logr.Get().Info("Equipment deleted successfully")

	return DeleteEquipmentResp{}, nil
}
