package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAttachmentsByEquipmentIDQuery struct {
	EquipmentID int `json:"equipment_id"`
	Read        ports.EquipmentAttachmentRead
}

type GetAttachmentsByEquipmentIDResp struct {
	Attachments []equipment.Attachment
}

func (qry *GetAttachmentsByEquipmentIDQuery) Handle(ctx context.Context) (any, error) {
	attachments, err := qry.Read.GetByEquipmentID(ctx, qry.EquipmentID)
	if err != nil {
		logr.Get().Errorf("failed to get attachments: %v", err)
		return GetAttachmentsByEquipmentIDResp{}, fmt.Errorf("failed to get attachments: %w", err)
	}

	return GetAttachmentsByEquipmentIDResp{Attachments: attachments}, nil
}
