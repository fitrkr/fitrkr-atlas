package muscle

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrEmptyMuscleGroup   = errors.New("empty muscle group")
	ErrInvalidMuscleGroup = errors.New("invalid muscle group")
)

type MuscleGroup string

const (
	CHEST     MuscleGroup = "chest"
	BACK      MuscleGroup = "back"
	SHOULDERS MuscleGroup = "shoulders"
	ARMS      MuscleGroup = "arms"
	LEGS      MuscleGroup = "legs"
	CORE      MuscleGroup = "core"
	NECK      MuscleGroup = "neck"
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

func NewMuscleGroupType(name string) (MuscleGroup, error) {
	if name == "" {
		return "", ErrEmptyMuscleGroup
	}

	name = strings.ToLower(name)
	switch name {
	case "chest", "back", "shoulders", "arms", "legs", "core", "neck":
		return MuscleGroup(name), nil
	default:
		return "", ErrInvalidMuscleGroup
	}
}

func (m *Group) Touch() {
	m.UpdatedAt = time.Now()
}
