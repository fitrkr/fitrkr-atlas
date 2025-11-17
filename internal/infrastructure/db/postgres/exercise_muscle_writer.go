package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"
	"github.com/lib/pq"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type ExerciseMuscleWriter struct {
	db *sql.DB
}

func newExerciseMuscleWriter(db *sql.DB) *ExerciseMuscleWriter {
	return &ExerciseMuscleWriter{db: db}
}

const CreateExerciseMuscles = `INSERT INTO exercise_muscle (exercise_id, muscle_id, activation, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

func (w *ExerciseMuscleWriter) Add(ctx context.Context, muscles []exercise.ExerciseMuscle) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		stmt, err := tx.PrepareContext(ctx, CreateExerciseMuscles)
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, muscle := range muscles {
			_, err := stmt.ExecContext(ctx,
				muscle.ExerciseID,
				muscle.MuscleID,
				muscle.Activation,
				muscle.CreatedAt,
				muscle.UpdatedAt,
			)
			if err != nil {
				return err
			}
		}

		logr.Get().Infof("Created %d muscles!", len(muscles))
		return nil
	})
}

const DeleteExerciseMuscles = `DELETE FROM exercise_muscle WHERE id = ANY($1)`

func (w *ExerciseMuscleWriter) Delete(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteExerciseMuscles, pq.Array(ids))
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrExerciseMuscleNotFound
		}

		logr.Get().Infof("Deleted %d muscles!", rowsAffected)
		return nil
	})
}
