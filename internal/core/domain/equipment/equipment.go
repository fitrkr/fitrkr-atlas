// Package equipment
package equipment

import (
	"errors"
	"strings"
	"time"
)

var ErrEmptyEquipmentName = errors.New("empty equipment name")

type Equipment struct {
	ID          *int      `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func New(name, description string) (Equipment, error) {
	if name == "" {
		return Equipment{}, ErrEmptyEquipmentName
	}

	return Equipment{
		Name:        strings.ToLower(name),
		Description: &description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (e *Equipment) Touch() {
	e.UpdatedAt = time.Now()
}
