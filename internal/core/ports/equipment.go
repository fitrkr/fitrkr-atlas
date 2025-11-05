package ports

import (
	"context"
	"errors"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
)

var (
	ErrEquipmentNotFound  = errors.New("equipment does not exist")
	ErrAttachmentNotFound = errors.New("attachment does not exist")
)

type EquipmentWrite interface {
	Add(ctx context.Context, eq equipment.Equipment) error
	Update(ctx context.Context, eq equipment.Equipment) error
	Delete(ctx context.Context, id int) error
}

type EquipmentRead interface {
	GetByID(ctx context.Context, id int) (*equipment.Equipment, error)
	GetAll(ctx context.Context) ([]equipment.Equipment, error)
}

type EquipmentAttachmentWrite interface {
	Add(ctx context.Context, at equipment.Attachment) error
	Update(ctx context.Context, at equipment.Attachment) error
	Delete(ctx context.Context, id int) error
}

type EquipmentAttachmentRead interface {
	GetByID(ctx context.Context, id int) (*equipment.Attachment, error)
	GetAll(ctx context.Context) ([]equipment.Attachment, error)
	GetByEquipmentID(ctx context.Context, equipmentID int) ([]equipment.Attachment, error)
}
