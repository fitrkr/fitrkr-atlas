package exercise

import (
	"errors"
	"time"
)

var ErrEmptySubcategoryID = errors.New("empty subcategory id")

type ExerciseCategory struct {
	ID            *int
	ExerciseID    int
	SubcategoryID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewExerciseCategory(exerciseID, subcategoryID int) (ExerciseCategory, error) {
	if exerciseID < 0 {
		return ExerciseCategory{}, ErrEmptyExericiseID
	}
	if subcategoryID < 0 {
		return ExerciseCategory{}, ErrEmptySubcategoryID
	}
	return ExerciseCategory{
		ExerciseID:    exerciseID,
		SubcategoryID: subcategoryID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}
