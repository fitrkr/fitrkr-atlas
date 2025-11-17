package muscle

import (
	"errors"
	"strings"
)

var ErrInvalidMuscleGroup = errors.New("invalid muscle group")

type MuscleGroup int

const (
	ARMS MuscleGroup = iota + 1
	BACK
	CHEST
	CORE
	LEGS
	NECK
	SHOULDERS
)

func NewMuscleGroupType(muscleGroup string) (MuscleGroup, error) {
	muscleGroup = strings.TrimSpace(strings.ToLower(muscleGroup))
	switch muscleGroup {
	case "arms":
		return ARMS, nil
	case "back":
		return BACK, nil
	case "chest":
		return CHEST, nil
	case "core":
		return CORE, nil
	case "legs":
		return LEGS, nil
	case "neck":
		return NECK, nil
	case "shoulders":
		return SHOULDERS, nil
	default:
		return 0, ErrInvalidMuscleGroup
	}
}

func (m MuscleGroup) ToString() string {
	switch m {
	case ARMS:
		return "arms"
	case BACK:
		return "back"
	case CHEST:
		return "chest"
	case CORE:
		return "core"
	case LEGS:
		return "legs"
	case NECK:
		return "neck"
	case SHOULDERS:
		return "shoulders"
	default:
		return ""
	}
}
