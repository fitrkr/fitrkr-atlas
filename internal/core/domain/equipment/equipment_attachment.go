package equipment

import "time"

type EquipmentAttachment struct {
	ID           *int
	EquipmentID  int
	AttachmentID int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewEquipmentAttachment(equipmentID, attachmentID int) EquipmentAttachment {
	return EquipmentAttachment{
		EquipmentID:  equipmentID,
		AttachmentID: attachmentID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
