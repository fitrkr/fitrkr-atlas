package exercise

import (
	"errors"
	"strings"
)

type Position int

const (
	HALFKNEELING Position = iota + 1
	HANGING
	INVERTED
	KNEELING
	PRONE
	QUADRUPED
	SIDELYING
	SITTING
	STANDING
	SUPINE
)

var ErrInvalidPosition = errors.New("invalid body position")

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
		return 0, ErrInvalidPosition
	}
}

func (p Position) ToString() string {
	switch p {
	case HALFKNEELING:
		return "half-kneeling"
	case HANGING:
		return "hanging"
	case INVERTED:
		return "inverted"
	case KNEELING:
		return "kneeling"
	case PRONE:
		return "prone"
	case QUADRUPED:
		return "quadruped"
	case SIDELYING:
		return "side-lying"
	case SITTING:
		return "sitting"
	case STANDING:
		return "standing"
	case SUPINE:
		return "supine"
	default:
		return ""
	}
}
