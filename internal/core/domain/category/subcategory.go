package category

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrEmptySubCategory  = errors.New("empty subcategory")
	ErrInvalidCategoryID = errors.New("invalid category id")
)

type Subcategory struct {
	ID         *int
	CategoryID int
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewSubcategory(name string, categoryID int) (Subcategory, error) {
	if name == "" {
		return Subcategory{}, ErrEmptySubCategory
	}
	if categoryID < 0 {
		return Subcategory{}, ErrInvalidCategoryID
	}
	return Subcategory{
		Name:       strings.TrimSpace(strings.ToLower(name)),
		CategoryID: categoryID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}

func (sc *Subcategory) Touch() {
	sc.UpdatedAt = time.Now()
}
