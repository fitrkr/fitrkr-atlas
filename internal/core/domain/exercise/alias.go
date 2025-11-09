package exercise

import (
	"errors"
	"strings"
	"time"
)

var ErrEmptyAliasName = errors.New("empty alias name")

type Alias struct {
	ID           *int
	ExerciseID   int
	Name         string
	LanguageCode string // TODO Make a type for it
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewAlias(exerciseID int, name string, languageCode string) (Alias, error) {
	if exerciseID < 0 {
		return Alias{}, ErrEmptyExericiseID
	}
	if name == "" {
		return Alias{}, ErrEmptyAliasName
	}
	return Alias{
		ExerciseID:   exerciseID,
		Name:         strings.TrimSpace(strings.ToLower(name)),
		LanguageCode: languageCode,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}
