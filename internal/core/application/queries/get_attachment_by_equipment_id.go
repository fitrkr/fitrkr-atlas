package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/equipment"
)

type GetAttachmentByEquipmentIDQuery struct {
	EquipmentID int `json:"id"`
}

type GetAttachmentByEquipmentIDResp struct {
	EquipmentAttachment []*equipment.EquipmentAttachment
}

func (qry *GetAttachmentByEquipmentIDQuery) Handle(ctx context.Context) (any, error) {
	equipmentAttachment, err := read.Equipment.Attachment.GetByEquipmentID(ctx, qry.EquipmentID)
	if err != nil {
		return GetAttachmentByEquipmentIDResp{}, fmt.Errorf("failed to get equipment attachment: %w", err)
	}

	return GetAttachmentByEquipmentIDResp{EquipmentAttachment: equipmentAttachment}, nil
}

func init() {
	register(&GetAttachmentByEquipmentIDQuery{})
}
