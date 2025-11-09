package exercise

import (
	"errors"
	"strings"
)

var ErrInvalidActivationLevel = errors.New("invalid activation level")

type ActivationLevel string

const (
	PRIMARY   ActivationLevel = "primary"
	SECONDARY ActivationLevel = "secondary"
	TERTIARY  ActivationLevel = "tertiary"
)

func NewActivationLevel(level string) (ActivationLevel, error) {
	switch strings.TrimSpace(strings.ToLower(level)) {
	case "primary", "secondary", "tertiary":
		return ActivationLevel(level), nil
	default:
		return "", ErrInvalidActivationLevel
	}
}
