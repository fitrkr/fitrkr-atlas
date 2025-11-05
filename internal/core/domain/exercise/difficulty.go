package exercise

import (
	"errors"
	"strings"
)

type Difficulty int

const (
	DifficultyBeginner Difficulty = iota
	DifficultyIntermediate
	DifficultyAdvanced
	DifficultyElite
)

var ErrInvalidDifficulty = errors.New("invalid difficulty")

func NewDifficulty(difficulty string) (Difficulty, error) {
	switch strings.ToLower(difficulty) {
	case "beginner":
		return DifficultyBeginner, nil
	case "intermediate":
		return DifficultyIntermediate, nil
	case "advanced":
		return DifficultyAdvanced, nil
	case "elite":
		return DifficultyElite, nil
	default:
		return DifficultyBeginner, ErrInvalidDifficulty
	}
}

func (d Difficulty) ToString() string {
	switch d {
	case DifficultyBeginner:
		return "Beginner"
	case DifficultyIntermediate:
		return "Intermediate"
	case DifficultyAdvanced:
		return "Advanced"
	case DifficultyElite:
		return "Elite"
	default:
		return "Unknown"
	}
}
