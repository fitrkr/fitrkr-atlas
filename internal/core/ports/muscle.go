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

type MuscleGroupWrite interface {
	Add(ctx context.Context, muscle muscle.Group) error
	Update(ctx context.Context, muscle muscle.Group) error
	Delete(ctx context.Context, id int) error
}

type MuscleGroupRead interface {
	GetByID(ctx context.Context, id int) (*muscle.Group, error)
	GetAll(ctx context.Context) ([]muscle.Group, error)
}

type MuscleWrite interface {
	Add(ctx context.Context, muscle muscle.Muscle) error
	Update(ctx context.Context, muscle muscle.Muscle) error
	Delete(ctx context.Context, id int) error
}

type MuscleRead interface {
	GetByID(ctx context.Context, id int) (*muscle.Muscle, error)
	GetAll(ctx context.Context) ([]muscle.Muscle, error)
	GetByMuscleGroupID(ctx context.Context, muscleGroupID int) ([]muscle.Muscle, error)
}
