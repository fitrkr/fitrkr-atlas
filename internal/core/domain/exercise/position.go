package exercise

import (
	"errors"
	"strings"
)

type Position int

const (
	STANDING Position = iota + 1
	SITTING
	KNEELING
	PRONE
	SUPINE
	SIDELYING
	QUADRUPED
	HALFKNEELING
	INVERTED
	HANGING
)

var ErrInvalidBodyPosition = errors.New("invalid body position")

func NewBodyPosition(position string) (Position, error) {
	switch strings.TrimSpace(strings.ToLower(position)) {
	case "standing":
		return STANDING, nil
	case "sitting":
		return SITTING, nil
	case "kneeling":
		return KNEELING, nil
	case "prone":
		return PRONE, nil
	case "supine":
		return SUPINE, nil
	case "side-lying":
		return SIDELYING, nil
	case "quadruped":
		return QUADRUPED, nil
	case "half-kneeling":
		return HALFKNEELING, nil
	case "inverted":
		return INVERTED, nil
	case "hanging":
		return HANGING, nil
	default:
		return 0, ErrInvalidBodyPosition
	}
}

func (p Position) ToString() string {
	switch p {
	case STANDING:
		return "standing"
	case SITTING:
		return "sitting"
	case KNEELING:
		return "kneeling"
	case PRONE:
		return "prone"
	case SUPINE:
		return "supine"
	case SIDELYING:
		return "side-lying"
	case QUADRUPED:
		return "quadruped"
	case HALFKNEELING:
		return "half-kneeling"
	case INVERTED:
		return "inverted"
	case HANGING:
		return "hanging"
	default:
		return ""
	}
}
