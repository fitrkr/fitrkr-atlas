package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"
	"github.com/lib/pq"

	"github.com/fitrkr/atlas/internal/core/domain/exercise"
	"github.com/fitrkr/atlas/internal/core/ports"
)

type ExerciseCategoryWriter struct {
	db *sql.DB
}

func newExerciseCategoryWriter(db *sql.DB) *ExerciseCategoryWriter {
	return &ExerciseCategoryWriter{db: db}
}

const CreateExerciseCategories = `INSERT INTO exercise_category (exercise_id, category_id, created_at, updated_at) VALUES ($1, $2, $3, $4)`

func (w *ExerciseCategoryWriter) Add(ctx context.Context, categories []exercise.ExerciseCategory) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		stmt, err := tx.PrepareContext(ctx, CreateExerciseCategories)
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, category := range categories {
			_, err := stmt.ExecContext(ctx,
				category.ExerciseID,
				category.CategoryID,
				category.CreatedAt,
				category.UpdatedAt,
			)
			if err != nil {
				return err
			}
		}

		logr.Get().Infof("Created %d categories!", len(categories))
		return nil
	})
}

const DeleteExerciseCategories = `DELETE FROM exercise_category WHERE id = ANY($1)`

func (w *ExerciseCategoryWriter) Delete(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteExerciseCategories, pq.Array(ids))
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrExerciseCategoryNotFound
		}

		logr.Get().Infof("Deleted %d categories!", rowsAffected)
		return nil
	})
}
