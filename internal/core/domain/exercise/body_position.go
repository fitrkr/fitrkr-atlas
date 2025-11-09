package exercise

import (
	"errors"
	"strings"
)

type BodyPosition string

const (
	STANDING     BodyPosition = "standing"
	SITTING      BodyPosition = "sitting"
	KNEELING     BodyPosition = "kneeling"
	PRONE        BodyPosition = "prone"
	SUPINE       BodyPosition = "supine"
	SIDELYING    BodyPosition = "side-lying"
	QUADRUPED    BodyPosition = "quadruped"
	HALFKNEELING BodyPosition = "half-kneeling"
	INVERTED     BodyPosition = "inverted"
	HANGING      BodyPosition = "hanging"
)

var ErrInvalidBodyPosition = errors.New("invalid body position")

func NewBodyPosition(position string) (BodyPosition, error) {
	switch strings.TrimSpace(strings.ToLower(position)) {
	case "standing", "sitting", "kneeling", "prone", "supine", "side-lying", "quadruped", "half-kneeling", "inverted", "hanging":
		return BodyPosition(position), nil
	// TODO for positions with a dash add logic here for it!
	default:
		return "", ErrInvalidBodyPosition
	}
}
