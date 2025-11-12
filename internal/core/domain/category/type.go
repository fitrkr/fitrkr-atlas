package category

import "strings"

type CategoryType int

const (
	STRENGTH CategoryType = iota + 1
	CARDIO
	FLEXIBILITY
)

func NewCategoryType(categoryType string) (CategoryType, error) {
	if categoryType == "" {
		return 0, ErrEmptyCategory
	}

	categoryType = strings.TrimSpace(strings.ToLower(categoryType))
	switch categoryType {
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
	case STRENGTH:
		return "strength"
	case CARDIO:
		return "cardio"
	case FLEXIBILITY:
		return "flexibility"

	default:
		return ""
	}
}
