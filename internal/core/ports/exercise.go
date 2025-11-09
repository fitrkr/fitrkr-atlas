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

	ErrExerciseMediaNotFound      = errors.New("exercise media does not exist")
	ErrExerciseMediaPrimaryExists = errors.New("exercise media primiray already exist")
	ErrExerciseMediaOrderExists   = errors.New("exercise media order already exist")
	ErrDuplicateExerciseMedia     = errors.New("exercise media already exists")

	ErrExerciseInstructionNotFound    = errors.New("exercise instruction does not exist")
	ErrExerciseInstructionOrderExists = errors.New("exercise instruction order already exist")
	ErrDuplicateExerciseInstruction   = errors.New("exercise instruction already exists")
)

type ExerciseWrite interface {
	Add(ctx context.Context, exercise exercise.Exercise) (int, error) // return the id
	Update(ctx context.Context, exercise exercise.Exercise) error
	Delete(ctx context.Context, id int) error
}

type ExerciseRead interface {
	GetByID(ctx context.Context, id int) (*exercise.Exercise, error)
	GetByName(ctx context.Context, name string) (*exercise.Exercise, error)
	// TODO GetAll will be different since it will use the denormalized table
	// will likely not be from here
}

type ExerciseAliasWrite interface {
	Add(ctx context.Context, alias exercise.Alias) error
	Update(ctx context.Context, alias exercise.Alias) error
	Delete(ctx context.Context, id int) error
}

type ExerciseAliasRead interface {
	GetByID(ctx context.Context, id int) (*exercise.Alias, error)
	GetByName(ctx context.Context, name string) (*exercise.Alias, error)
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
	Update(ctx context.Context, muscle exercise.ExerciseMuscle) error // TODO check how update would work for this
	Delete(ctx context.Context, id int) error
}

type ExerciseMuscleRead interface {
	GetByID(ctx context.Context, id int) (*exercise.ExerciseMuscle, error)
	GetByExerciseAndMuscleID(ctx context.Context, exerciseID, muscleID int) (*exercise.ExerciseMuscle, error)
}

type ExerciseCategoryWrite interface {
	Add(ctx context.Context, exerciseCategory exercise.ExerciseCategory) error
	Delete(ctx context.Context, id int) error
}

type ExerciseCategoryRead interface {
	GetByID(ctx context.Context, id int) (*exercise.ExerciseCategory, error)
	GetByExerciseAndCategoryID(ctx context.Context, exerciseID, categoryID int) (*exercise.ExerciseCategory, error)
}

type ExerciseMediaWrite interface {
	Add(ctx context.Context, media exercise.Media) error
	Update(ctx context.Context, media exercise.Media) error
	Delete(ctx context.Context, id int) error
}

type ExerciseMediaRead interface {
	GetByID(ctx context.Context, id int) (*exercise.Media, error)
	GetPrimaryByExerciseID(ctx context.Context, exerciseID int) (*exercise.Media, error)
	GetOrderByExerciseID(ctx context.Context, id, order int) (*exercise.Media, error)
}

type ExerciseInstructionWrite interface {
	Add(ctx context.Context, instruction exercise.Instruction) error
	Update(ctx context.Context, instruction exercise.Instruction) error
	Delete(ctx context.Context, id int) error
}

type ExerciseInstructionRead interface {
	GetByID(ctx context.Context, id int) (*exercise.Instruction, error)
	GetOrderByExerciseID(ctx context.Context, id, order int) (*exercise.Media, error)
}

// TODO implement this last

type ExerciseVariationWrite interface {
	Add(ctx context.Context, exerciseID, variationID int) error
	Update(ctx context.Context, exerciseID, variationID int) error
	Delete(ctx context.Context, exerciseID, variationID int) error
}

type ExerciseAlternateWrite interface {
	Add(ctx context.Context, exerciseID, alternateID int) error
	Update(ctx context.Context, exerciseID, alternateID int) error
	Delete(ctx context.Context, exerciseID, alternateID int) error
}

// TODO
// type ExerciseAlternateRead interface {}
// type ExerciseVariationRead interface {}
