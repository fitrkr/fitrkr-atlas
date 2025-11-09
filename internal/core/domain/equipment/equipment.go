// Package equipment
package equipment

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrEmptyEquipmentName   = errors.New("empty equipment name")
	ErrInvalidEquipmentType = errors.New("invalid equipment type")
)

type Equipment struct {
	ID                 *int          `json:"id"`
	Name               string        `json:"name"`
	Description        *string       `json:"description"`
	Type               EquipmentType `json:"type"`
	AllowedAttachments bool          `json:"allowed_attachments"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
}

func New(name, description string, equipmentType EquipmentType, allowedAttachments bool) (Equipment, error) {
	if name == "" {
		return Equipment{}, ErrEmptyEquipmentName
	}

	return Equipment{
		Name:               strings.TrimSpace(strings.ToLower(name)),
		Description:        &description,
		Type:               equipmentType,
		AllowedAttachments: allowedAttachments,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}, nil
}

func (e *Equipment) Touch() {
	e.UpdatedAt = time.Now()
}
