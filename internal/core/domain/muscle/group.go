package muscle

import (
	"errors"
	"strings"
)

var (
	ErrInvalidMuscleGroupType = errors.New("invalid muscle group type")
	ErrEmptyMuscleGroup       = errors.New("empty muscle group")
)

type MuscleGroup int

const (
	CHEST MuscleGroup = iota + 1
	BACK
	SHOULDERS
	ARMS
	LEGS
	CORE
	NECK
)

func NewMuscleGroupType(groupType string) (MuscleGroup, error) {
	if groupType == "" {
		return 0, ErrEmptyMuscleGroup
	}

	groupType = strings.TrimSpace(strings.ToLower(groupType))
	switch groupType {
	case "chest":
		return CHEST, nil
	case "back":
		return BACK, nil
	case "shoulders":
		return SHOULDERS, nil
	case "arms":
		return ARMS, nil
	case "legs":
		return LEGS, nil
	case "core":
		return CORE, nil
	case "neck":
		return NECK, nil
	default:
		return 0, ErrInvalidMuscleGroupType
	}
}

func (m MuscleGroup) ToString() string {
	switch m {
	case CHEST:
		return ""
	case BACK:
		return "back"
	case SHOULDERS:
		return "shoulders"
	case ARMS:
		return "arms"
	case LEGS:
		return "legs"
	case CORE:
		return "core"
	case NECK:
		return "neck"
	default:
		return ""
	}
}
