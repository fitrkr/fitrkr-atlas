package ports

import (
	"context"
	"errors"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
)

var (
	ErrMuscleGroupNotFound = errors.New("muscle group does not exist")
	ErrMuscleNotFound      = errors.New("muscle does not exist")
)

type MuscleWrite interface {
	Add(ctx context.Context, muscle muscle.Muscle) error
	Update(ctx context.Context, muscle muscle.Muscle) error
	Delete(ctx context.Context, id int) error
}

type MuscleRead interface {
	GetByID(ctx context.Context, id int) (*muscle.Muscle, error)
	GetAll(ctx context.Context) ([]*muscle.Muscle, error)
	GetByGroupType(ctx context.Context, muscleType string) ([]*muscle.Muscle, error)
}
