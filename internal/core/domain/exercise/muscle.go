package exercise

import (
	"errors"
	"time"
)

var ErrEmptyMuscleID = errors.New("empty muscle id")

type ExerciseMuscle struct {
	ID         *int       `json:"id"`
	ExerciseID int        `json:"exercise_id"`
	MuscleID   int        `json:"muscle_id"`
	Activation Activation `json:"activation"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func NewExerciseMuscle(exerciseID, muscleID int, activation Activation) (ExerciseMuscle, error) {
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
