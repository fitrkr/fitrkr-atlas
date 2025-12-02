// Package ports

package ports

import (
	"context"
	"errors"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/exercise"
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

	ErrExerciseAttachmentNotFound  = errors.New("exercise attachment does not exist")
	ErrDuplicateExerciseAttachment = errors.New("exercise attachment already exists")

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

type ExerciseAttachmentWrite interface {
	Add(ctx context.Context, attachments []exercise.ExerciseAttachment) error
	Delete(ctx context.Context, ids []int) error
}

type ExerciseAttachmentRead interface {
	GetByID(ctx context.Context, id int) (*exercise.ExerciseAttachment, error)
	GetByExerciseID(ctx context.Context, exerciseID int) ([]*exercise.ExerciseAttachment, error)
}

type ExerciseMuscleWrite interface {
	Add(ctx context.Context, muscles []exercise.ExerciseMuscle) error
	Delete(ctx context.Context, ids []int) error
}

type ExerciseMuscleRead interface {
	GetByID(ctx context.Context, id int) (*exercise.ExerciseMuscle, error)
	GetByExerciseID(ctx context.Context, exerciseID int) ([]*exercise.ExerciseMuscle, error)
}

type ExerciseCategoryWrite interface {
	Add(ctx context.Context, exerciseCategories []exercise.ExerciseCategory) error
	Delete(ctx context.Context, ids []int) error
}

type ExerciseCategoryRead interface {
	GetByID(ctx context.Context, id int) (*exercise.ExerciseCategory, error)
	GetByExerciseID(ctx context.Context, exerciseID int) ([]*exercise.ExerciseCategory, error)
}

type ExerciseAliasWrite interface {
	Add(ctx context.Context, aliases []exercise.Alias) error
	Delete(ctx context.Context, ids []int) error
}

type ExerciseAliasRead interface {
	GetByID(ctx context.Context, id int) (*exercise.Alias, error)
	GetByName(ctx context.Context, name string) (*exercise.Alias, error)
	GetByExerciseID(ctx context.Context, exerciseID int) ([]*exercise.Alias, error)
}
