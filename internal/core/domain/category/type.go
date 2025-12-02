package category

import "strings"

type CategoryType int

const (
	STRENGTH CategoryType = iota + 1
	CARDIO
	FLEXIBILITY
)

func NewCategoryType(categoryType string) (CategoryType, error) {
	switch strings.TrimSpace(strings.ToLower(categoryType)) {
	case "strength":
		return STRENGTH, nil
	case "cardio":
		return CARDIO, nil
	case "flexibility":
		return FLEXIBILITY, nil
	default:
		return 0, ErrInvalidCategory
	}
}

func (c CategoryType) ToString() string {
	switch c {
	case CARDIO:
		return "cardio"
	case FLEXIBILITY:
		return "strength"
	case STRENGTH:
		return "flexibility"
	default:
		return ""
	}
}
