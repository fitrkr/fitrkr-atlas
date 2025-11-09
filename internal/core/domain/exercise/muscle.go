package exercise

import (
	"errors"
	"time"
)

var ErrEmptyMuscleID = errors.New("empty muscle id")

type ExerciseMuscle struct {
	ID              *int
	ExerciseID      int
	MuscleID        int
	ActivationLevel ActivationLevel
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewExerciseMuscle(exerciseID, muscleID int, activationLevel ActivationLevel) (ExerciseMuscle, error) {
	if exerciseID < 0 {
		return ExerciseMuscle{}, ErrEmptyExericiseID
	}
	if muscleID < 0 {
		return ExerciseMuscle{}, ErrEmptyMuscleID
	}
	return ExerciseMuscle{
		ExerciseID:      exerciseID,
		MuscleID:        muscleID,
		ActivationLevel: activationLevel,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, nil
}
