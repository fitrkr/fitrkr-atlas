package muscle

import "strings"

type MuscleGroup string

const (
	CHEST     MuscleGroup = "chest"
	BACK      MuscleGroup = "back"
	SHOULDERS MuscleGroup = "shoulders"
	ARMS      MuscleGroup = "arms"
	LEGS      MuscleGroup = "legs"
	CORE      MuscleGroup = "core"
	NECK      MuscleGroup = "neck"
)

func NewMuscleGroupType(name string) (MuscleGroup, error) {
	if name == "" {
		return "", ErrEmptyMuscleGroup
	}

	name = strings.TrimSpace(strings.ToLower(name))
	switch name {
	case "chest", "back", "shoulders", "arms", "legs", "core", "neck":
		return MuscleGroup(name), nil
	default:
		return "", ErrInvalidMuscleGroup
	}
}
