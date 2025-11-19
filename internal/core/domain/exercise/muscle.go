package exercise

import (
	"errors"
	"time"
)

var ErrEmptyMuscleID = errors.New("empty muscle id")

type ExerciseMuscle struct {
	ID         *int
	ExerciseID int
	MuscleID   int
	Activation string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewExerciseMuscle(exerciseID, muscleID int, activation string) (ExerciseMuscle, error) {
	if exerciseID < 0 {
		return ExerciseMuscle{}, ErrEmptyExericiseID
	}
	if muscleID < 0 {
		return ExerciseMuscle{}, ErrEmptyMuscleID
	}
	return ExerciseMuscle{
		ExerciseID: exerciseID,
		MuscleID:   muscleID,
		Activation: activation,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
