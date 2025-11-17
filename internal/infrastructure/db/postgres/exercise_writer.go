package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type ExerciseWriter struct {
	db *sql.DB
}

func newExerciseWriter(db *sql.DB) *ExerciseWriter {
	return &ExerciseWriter{db: db}
}

const CreateExercise = `
	INSERT INTO exercise (name, description, difficulty, position, equipment_id, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id
`

func (w *ExerciseWriter) Add(ctx context.Context, ex exercise.Exercise) (exercise.Exercise, error) {
	err := WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		// Execute insert and scan the returned ID
		err := tx.QueryRowContext(
			ctx,
			CreateExercise,
			ex.Name,
			ex.Description,
			ex.Difficulty,
			ex.Position,
			ex.EquipmentID,
			ex.CreatedAt,
			ex.UpdatedAt,
		).Scan(&ex.ID)
		if err != nil {
			return err
		}

		logr.Get().Info("Exercise created!")
		return nil
	})
	if err != nil {
		return exercise.Exercise{}, err
	}

	return ex, nil
}

const UpdateExercise = `
	UPDATE exercise
	SET name = $2, 
		description = $3,
		difficulty = $4,
		position = $5,
		equipment_id = $6,
		updated_at = $7
	WHERE id = $1 AND deleted_at IS NULL
`

func (w *ExerciseWriter) Update(ctx context.Context, ex exercise.Exercise) (exercise.Exercise, error) {
	err := WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(
			ctx,
			UpdateExercise,
			ex.ID,
			ex.Name,
			ex.Description,
			ex.Difficulty,
			ex.Position,
			ex.EquipmentID,
			ex.UpdatedAt,
		)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrExerciseNotFound
		}

		logr.Get().Info("Exercise updated!")
		return nil
	})
	if err != nil {
		return exercise.Exercise{}, err
	}

	return ex, nil
}

const DeleteExercise = `
	UPDATE exercise 
	SET deleted_at = CURRENT_TIMESTAMP,
		purge_at = CURRENT_TIMESTAMP + INTERVAL '45 days'
	WHERE id = $1 AND deleted_at IS NULL
`

func (w *ExerciseWriter) Delete(ctx context.Context, id int) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteExercise, id)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrExerciseNotFound
		}

		logr.Get().Info("Exercise soft deleted!")
		return nil
	})
}
