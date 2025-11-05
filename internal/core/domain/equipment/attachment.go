package equipment

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrEmptyAttachment    = errors.New("empty attachment")
	ErrInvalidEquipmentID = errors.New("invalid equipment id")
)

type Attachment struct {
	ID          *int
	EquipmentID int
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewAttachment(name string, equipmentID int) (Attachment, error) {
	if name == "" {
		return Attachment{}, ErrEmptyAttachment
	}
	if equipmentID < 0 {
		return Attachment{}, ErrInvalidEquipmentID
	}
	return Attachment{Name: strings.ToLower(name), EquipmentID: equipmentID}, nil
}

func (a *Attachment) Touch() {
	a.UpdatedAt = time.Now()
}
