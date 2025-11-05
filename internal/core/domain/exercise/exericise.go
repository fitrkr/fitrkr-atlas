// Package exercise
package exercise

import (
	"errors"
	"time"
)

var ErrEmptyName = errors.New("empty name")

type Exercise struct {
	ID          *int       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Difficulty  Difficulty `json:"difficulty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	PurgedAt    *time.Time `json:"purged_at"`
	Media       []*Media
}

func New(name, description string, difficulty Difficulty) (Exercise, error) {
	if name == "" {
		return Exercise{}, ErrEmptyName
	}
	return Exercise{
		Name:        name,
		Description: description,
		Difficulty:  difficulty,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
