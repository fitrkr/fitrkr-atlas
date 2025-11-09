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
	Name      CategoryType
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(name CategoryType) (Category, error) {
	return Category{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (c *Category) Touch() {
	c.UpdatedAt = time.Now()
}
