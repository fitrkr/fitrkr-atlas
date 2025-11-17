package ports

import (
	"context"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/view"
)

type ViewWrite interface {
	Add(ctx context.Context, view view.View) (int, error)
	Update(ctx context.Context, view view.View) error
	Delete(ctx context.Context, id int) error
}

type ViewRead interface {
	GetByID(ctx context.Context, exerciseID int) (*view.View, error)
	GetAll(ctx context.Context) ([]*view.View, error)
}
