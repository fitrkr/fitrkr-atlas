package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"
	"github.com/lib/pq"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type ExerciseAliasWriter struct {
	db *sql.DB
}

func newExerciseAliasWriter(db *sql.DB) *ExerciseAliasWriter {
	return &ExerciseAliasWriter{db: db}
}

const CreateExerciseAlias = `INSERT INTO exercise_alias (exercise_id, name, language_code, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

func (w *ExerciseAliasWriter) Add(ctx context.Context, aliases []exercise.Alias) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		stmt, err := tx.PrepareContext(ctx, CreateExerciseAlias)
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, alias := range aliases {
			_, err := stmt.ExecContext(ctx,
				alias.ExerciseID,
				alias.Name,
				alias.LanguageCode,
				alias.CreatedAt,
				alias.UpdatedAt,
			)
			if err != nil {
				return err
			}
		}

		logr.Get().Infof("Created %d aliases!", len(aliases))
		return nil
	})
}

const DeleteExerciseAliases = `DELETE FROM exercise_alias WHERE id = ANY($1)`

func (w *ExerciseAliasWriter) Delete(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteExerciseAliases, pq.Array(ids))
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrExerciseAliasNotFound
		}

		logr.Get().Infof("Deleted %d aliases!", rowsAffected)
		return nil
	})
}
