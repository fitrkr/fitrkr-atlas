// Package ports
package ports

import (
	"context"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
)

type ExerciseWrite interface {
	Add(ctx context.Context, exercise exercise.Exercise) error
	Update(ctx context.Context, exercise exercise.Exercise) error
	Delete(ctx context.Context, id int) error
}

type ExerciseRead interface{}
