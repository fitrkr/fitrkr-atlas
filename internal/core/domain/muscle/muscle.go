// Package muscle
package muscle

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrEmptyMuscle          = errors.New("empty muscle")
	ErrInvalidMuscleGroupID = errors.New("invalid muscle group id")
)

type Muscle struct {
	ID            *int
	MuscleGroupID int
	Name          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func New(name string, muscleGroupID int) (Muscle, error) {
	if name == "" {
		return Muscle{}, ErrEmptyMuscle
	}
	if muscleGroupID < 0 {
		return Muscle{}, ErrInvalidMuscleGroupID
	}
	return Muscle{
		Name:          strings.TrimSpace(strings.ToLower(name)),
		MuscleGroupID: muscleGroupID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}

func (m *Muscle) Touch() {
	m.UpdatedAt = time.Now()
}
