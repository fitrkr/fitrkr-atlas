package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetEquipmentAttachmentByIDQuery struct {
	EquipmentID int `json:"id"`
	Read        ports.Read
}

type GetEquipmentAttachmentByIDResp struct {
	EquipmentAttachment []equipment.EquipmentAttachment
}

func (qry *GetEquipmentAttachmentByIDQuery) Handle(ctx context.Context) (any, error) {
	equipmentAttachment, err := qry.Read.Equipment.Attachment.GetByEquipmentID(ctx, qry.EquipmentID)
	if err != nil {
		return GetEquipmentAttachmentByIDResp{}, fmt.Errorf("failed to get equipment attachment: %w", err)
	}

	return GetEquipmentAttachmentByIDResp{EquipmentAttachment: equipmentAttachment}, nil
}
