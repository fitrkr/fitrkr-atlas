package exercise

import (
	"errors"
	"time"
)

var ErrEmptyCategoryID = errors.New("empty category id")

type ExerciseCategory struct {
	ID         *int      `json:"id"`
	ExerciseID int       `json:"exercise_id"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewExerciseCategory(exerciseID, categoryID int) (ExerciseCategory, error) {
	if exerciseID < 0 {
		return ExerciseCategory{}, ErrEmptyExericiseID
	}
	if categoryID < 0 {
		return ExerciseCategory{}, ErrEmptyCategoryID
	}
	return ExerciseCategory{
		ExerciseID: exerciseID,
		CategoryID: categoryID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
