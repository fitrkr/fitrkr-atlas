package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetEquipmentAttachmentByEquipmentIDQuery struct {
	EquipmentID int `json:"id"`
	Read        ports.Read
}

type GetEquipmentAttachmentByEquipmentIDResp struct {
	EquipmentAttachment []equipment.EquipmentAttachment
}

func (qry *GetEquipmentAttachmentByEquipmentIDQuery) Handle(ctx context.Context) (any, error) {
	equipmentAttachment, err := qry.Read.Equipment.Attachment.GetByEquipmentID(ctx, qry.EquipmentID)
	if err != nil {
		return GetEquipmentAttachmentByEquipmentIDResp{}, fmt.Errorf("failed to get equipment attachment: %w", err)
	}

	return GetEquipmentAttachmentByEquipmentIDResp{EquipmentAttachment: equipmentAttachment}, nil
}
