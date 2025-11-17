// Package category
package category

import (
	"errors"
	"time"
)

var (
	ErrEmptyCategory   = errors.New("empty category")
	ErrInvalidCategory = errors.New("invalid category")
)

type Category struct {
	ID        *int
	Name      string
	Type      CategoryType
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(name string, categoryType CategoryType) (Category, error) {
	if name == "" {
		return Category{}, ErrEmptyCategory
	}
	return Category{
		Name:      name,
		Type:      categoryType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (c *Category) Touch() {
	c.UpdatedAt = time.Now()
}
