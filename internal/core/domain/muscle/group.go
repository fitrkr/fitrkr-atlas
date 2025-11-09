package muscle

import (
	"errors"
	"time"
)

var (
	ErrEmptyMuscleGroup   = errors.New("empty muscle group")
	ErrInvalidMuscleGroup = errors.New("invalid muscle group")
)

type Group struct {
	ID          *int
	Name        MuscleGroup
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewGroup(name MuscleGroup, description string) (Group, error) {
	return Group{
		Name:        name,
		Description: &description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *Group) Touch() {
	m.UpdatedAt = time.Now()
}
