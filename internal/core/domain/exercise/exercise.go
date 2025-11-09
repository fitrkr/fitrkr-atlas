// Package exercise
package exercise

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrEmptyName           = errors.New("empty name")
	ErrEmptyExericiseID    = errors.New("empty exercise id")
	ErrEmptyEquipment      = errors.New("empty equipment")
	ErrInvalidExerciseType = errors.New("invalid exercise type")
)

type Exercise struct {
	ID           *int                 `json:"id"`
	Name         string               `json:"name"`
	Description  *string              `json:"description"`
	Difficulty   Difficulty           `json:"difficulty"`
	BodyPosition BodyPosition         `json:"body_position"`
	Type         ExerciseType         `json:"type"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
	DeletedAt    *time.Time           `json:"deleted_at"`
	PurgedAt     *time.Time           `json:"purged_at"`
	Equipment    []*ExerciseEquipment `json:"equipment"`
	Muscle       []*ExerciseMuscle    `json:"muscle"`
	Category     []*ExerciseCategory  `json:"category"`
	Alias        []*Alias             `json:"alias"`
	Media        []*Media             `json:"media"`
}

func New(name, description string, difficulty Difficulty, bodyPosition BodyPosition, exerciseType ExerciseType) (Exercise, error) {
	if name == "" {
		return Exercise{}, ErrEmptyName
	}
	return Exercise{
		Name:         strings.TrimSpace(strings.ToLower(name)),
		Description:  &description,
		Difficulty:   difficulty,
		BodyPosition: bodyPosition,
		Type:         exerciseType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

func (e *Exercise) Touch() {
	e.UpdatedAt = time.Now()
}
