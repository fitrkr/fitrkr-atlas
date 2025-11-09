package exercise

import (
	"errors"
	"strings"
)

var ErrInvalidActivationLevel = errors.New("invalid activation level")

type ActivationLevel int

const (
	PRIMARY = iota + 1
	SECONDARY
	TERTIARY
)

func NewActivationLevel(level string) (ActivationLevel, error) {
	switch strings.TrimSpace(strings.ToLower(level)) {
	case "primary":
		return PRIMARY, nil
	case "secondary":
		return SECONDARY, nil
	case "tertiary":
		return TERTIARY, nil
	default:
		return 0, ErrInvalidActivationLevel
	}
}

func (a ActivationLevel) ToString() string {
	switch a {
	case PRIMARY:
		return "primary"
	case SECONDARY:
		return "secondary"
	case TERTIARY:
		return "tertiary"
	default:
		return ""
	}
}
