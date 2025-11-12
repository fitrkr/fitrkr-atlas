// Package ports
package ports

import (
	"context"
	"errors"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
)

var (
	ErrExerciseNotFound  = errors.New("exercise does not exist")
	ErrDuplicateExercise = errors.New("exercise already exists")

	ErrExerciseAliasNotFound  = errors.New("exercise alias does not exist")
	ErrDuplicateExerciseAlias = errors.New("exercise alias already exists")

	ErrExerciseMuscleNotFound  = errors.New("exercise muscle does not exist")
	ErrDuplicateExerciseMuscle = errors.New("exercise muscle already exists")

	ErrExerciseEquipmentNotFound  = errors.New("exercise equipment does not exist")
	ErrDuplicateExerciseEquipment = errors.New("exercise equipment already exists")

	ErrExerciseCategoryNotFound  = errors.New("exercise category does not exist")
	ErrDuplicateExerciseCategory = errors.New("exercise category already exists")
)

type ExerciseWrite interface {
	Add(ctx context.Context, exercise exercise.Exercise) (exercise.Exercise, error)
	Update(ctx context.Context, exercise exercise.Exercise) (exercise.Exercise, error)
	Delete(ctx context.Context, id int) error
}

type ExerciseRead interface {
	GetByID(ctx context.Context, id int) (*exercise.Exercise, error)
	GetByName(ctx context.Context, name string) (*exercise.Exercise, error)
}

type ExerciseEquipmentWrite interface {
	Add(ctx context.Context, equipment exercise.ExerciseEquipment) error
	Delete(ctx context.Context, id int) error
}

type ExerciseEquipmentRead interface {
	GetByID(ctx context.Context, exerciseID int) (*exercise.ExerciseEquipment, error)
}

type ExerciseMuscleWrite interface {
	Add(ctx context.Context, muscle exercise.ExerciseMuscle) error
	Delete(ctx context.Context, id int) error
}

type ExerciseMuscleRead interface {
	GetByID(ctx context.Context, id int) (*exercise.ExerciseMuscle, error)
	GetByExerciseID(ctx context.Context, exerciseID int) ([]*exercise.ExerciseMuscle, error)
}

type ExerciseCategoryWrite interface {
	Add(ctx context.Context, exerciseCategory exercise.ExerciseCategory) error
	Delete(ctx context.Context, id int) error
}

type ExerciseCategoryRead interface {
	GetByID(ctx context.Context, id int) (*exercise.ExerciseCategory, error)
	GetByExerciseID(ctx context.Context, exerciseID int) ([]*exercise.ExerciseCategory, error)
}

type ExerciseAliasWrite interface {
	Add(ctx context.Context, alias exercise.Alias) error
	Delete(ctx context.Context, id int) error
}

type ExerciseAliasRead interface {
	GetByID(ctx context.Context, id int) (*exercise.Alias, error)
	GetByName(ctx context.Context, name string) (*exercise.Alias, error)
	GetByExerciseID(ctx context.Context, exerciseID int) ([]*exercise.Alias, error)
}
