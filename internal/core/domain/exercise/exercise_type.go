package exercise

import (
	"strings"
)

type ExerciseType int

const (
	BodyWeight ExerciseType = iota + 1
	FreeWeight
	Machine
)

func NewExerciseType(exerciseType string) (ExerciseType, error) {
	switch strings.TrimSpace(strings.ToLower(exerciseType)) {
	case "body_weight":
		return BodyWeight, nil
	case "free_weight":
		return FreeWeight, nil
	case "machine":
		return Machine, nil
	default:
		return 0, ErrInvalidExerciseType
	}
}

func (e ExerciseType) ToString() string {
	switch e {
	case BodyWeight:
		return "body_weight"
	case FreeWeight:
		return "free_weight"
	case Machine:
		return "machine"
	default:
		return ""
	}
}
