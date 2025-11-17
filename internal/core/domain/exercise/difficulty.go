package exercise

import (
	"errors"
	"strings"
)

type Difficulty int

const (
	BEGINNER Difficulty = iota + 1
	INTERMEDIATE
	ADVANCED
	ELITE
)

var ErrInvalidDifficulty = errors.New("invalid difficulty")

func NewDifficulty(difficulty string) (Difficulty, error) {
	switch strings.TrimSpace(strings.ToLower(difficulty)) {
	case "beginner":
		return BEGINNER, nil
	case "intermediate":
		return INTERMEDIATE, nil
	case "advanced":
		return ADVANCED, nil
	case "elite":
		return ELITE, nil
	default:
		return 0, ErrInvalidDifficulty
	}
}

func (d Difficulty) ToString() string {
	switch d {
	case BEGINNER:
		return "beginner"
	case INTERMEDIATE:
		return "intermediate"
	case ADVANCED:
		return "advanced"
	case ELITE:
		return "elite"
	default:
		return ""
	}
}
