package exercise

import (
	"errors"
	"time"
)

var ErrEmptyText = errors.New("empty text")

type Instruction struct {
	ID         *int
	ExerciseID int
	Text       string
	Order      int
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewInstruction(exerciseID, order int, text string) (Instruction, error) {
	if exerciseID < 0 {
		return Instruction{}, ErrEmptyExericiseID
	}
	if order < 1 {
		return Instruction{}, ErrEmptyOrder
	}
	if text == "" {
		return Instruction{}, ErrEmptyText
	}
	return Instruction{
		ExerciseID: exerciseID,
		Text:       text,
		Order:      order,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
