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
	ID          *int                  `json:"id"`
	Name        string                `json:"name"`
	Description *string               `json:"description"`
	Difficulty  string                `json:"difficulty"`
	Position    string                `json:"position"`
	EquipmentID *int                  `json:"equipment_id"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	DeletedAt   *time.Time            `json:"deleted_at"`
	PurgeAt     *time.Time            `json:"purge_at"`
	Attachment  []*ExerciseAttachment `json:"equipment"`
	Muscle      []*ExerciseMuscle     `json:"muscle"`
	Category    []*ExerciseCategory   `json:"category"`
	Alias       []*Alias              `json:"alias"`
}

func New(name, description, difficulty, position string) (Exercise, error) {
	if name == "" {
		return Exercise{}, ErrEmptyName
	}
	return Exercise{
		Name:        strings.TrimSpace(strings.ToLower(name)),
		Description: &description,
		Difficulty:  difficulty,
		Position:    position,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (e *Exercise) Touch() {
	e.UpdatedAt = time.Now()
}
