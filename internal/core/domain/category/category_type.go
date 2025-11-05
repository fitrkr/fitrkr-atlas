package category

import "strings"

type CategoryType string

const (
	STRENGTH    CategoryType = "strength"
	CARDIO      CategoryType = "cardio"
	FLEXIBILITY CategoryType = "flexibility"
)

func NewCategoryType(name string) (CategoryType, error) {
	if name == "" {
		return "", ErrEmptyCategory
	}

	name = strings.ToLower(name)
	switch name {
	case "strength", "cardio", "flexibility":
		return CategoryType(name), nil
	default:
		return "", ErrInvalidCategory
	}
}
