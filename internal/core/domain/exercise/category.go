package exercise

import (
	"errors"
	"time"
)

var ErrEmptySubcategoryID = errors.New("empty subcategory id")

type ExerciseCategory struct {
	ID         *int      `json:"id"`
	ExerciseID int       `json:"exercise_id"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewExerciseCategory(exerciseID, CategoryID int) (ExerciseCategory, error) {
	if exerciseID < 0 {
		return ExerciseCategory{}, ErrEmptyExericiseID
	}
	if CategoryID < 0 {
		return ExerciseCategory{}, ErrEmptySubcategoryID
	}
	return ExerciseCategory{
		ExerciseID: exerciseID,
		CategoryID: CategoryID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
