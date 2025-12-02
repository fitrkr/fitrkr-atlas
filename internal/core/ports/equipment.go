package ports

import (
	"context"
	"errors"

	"github.com/fitrkr/atlas/internal/core/domain/equipment"
)

var (
	ErrEquipmentNotFound           = errors.New("equipment does not exist")
	ErrAttachmentNotFound          = errors.New("attachment does not exist")
	ErrEquipmentAttachmentNotFound = errors.New("equipment attachment does not exist")
)

type EquipmentWrite interface {
	Add(ctx context.Context, equipment equipment.Equipment) (int, error)
	Update(ctx context.Context, equipment equipment.Equipment) error
	Delete(ctx context.Context, id int) error
}

type EquipmentRead interface {
	GetByID(ctx context.Context, id int) (*equipment.Equipment, error)
	GetAll(ctx context.Context) ([]*equipment.Equipment, error)
}

type AttachmentWrite interface {
	Add(ctx context.Context, attachment equipment.Attachment) error
	Update(ctx context.Context, attachment equipment.Attachment) error
	Delete(ctx context.Context, id int) error
}

type AttachmentRead interface {
	GetByID(ctx context.Context, id int) (*equipment.Attachment, error)
	GetAll(ctx context.Context) ([]*equipment.Attachment, error)
}

type EquipmentAttachmentWrite interface {
	Add(ctx context.Context, equipmentAttachment equipment.EquipmentAttachment) error
	Delete(ctx context.Context, id int) error
}

type EquipmentAttachmentRead interface {
	GetByID(ctx context.Context, id int) (*equipment.EquipmentAttachment, error)
	GetByEquipmentID(ctx context.Context, equipmentID int) ([]*equipment.EquipmentAttachment, error)
}
