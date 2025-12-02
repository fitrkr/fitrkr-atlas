package ports

import (
	"context"
	"errors"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/category"
)

var (
	ErrCategoryNotFound    = errors.New("category does not exist")
	ErrSubcategoryNotFound = errors.New("subcategory does not exist")
)

type CategoryWrite interface {
	Add(ctx context.Context, category category.Category) error
	Update(ctx context.Context, category category.Category) error
	Delete(ctx context.Context, id int) error
}

type CategoryRead interface {
	GetByID(ctx context.Context, id int) (*category.Category, error)
	GetAll(ctx context.Context) ([]*category.Category, error)
	GetByType(ctx context.Context, categoryType string) ([]*category.Category, error)
}
