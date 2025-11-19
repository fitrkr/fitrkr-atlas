// Package muscle
package muscle

import (
	"errors"
	"strings"
	"time"
)

var ErrEmptyMuscle = errors.New("empty muscle")

type Muscle struct {
	ID        *int
	Group     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(name, groupType string) (Muscle, error) {
	if name == "" {
		return Muscle{}, ErrEmptyMuscle
	}
	return Muscle{
		Name:      strings.TrimSpace(strings.ToLower(name)),
		Group:     groupType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (m *Muscle) Touch() {
	m.UpdatedAt = time.Now()
}
