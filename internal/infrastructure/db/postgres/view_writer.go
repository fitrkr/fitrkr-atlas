package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/cheezecakee/logr"
	"github.com/lib/pq"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/view"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type ViewWriter struct {
	db *sql.DB
}

func newViewWriter(db *sql.DB) *ViewWriter {
	return &ViewWriter{db: db}
}

const CreateView = `
	INSERT INTO exercise_view (
		id, name, description, difficulty, position, 
		alias, equipment, muscle, category,
		created_at, updated_at, deleted_at, purge_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING id
`

func (w *ViewWriter) Add(ctx context.Context, v view.View) (int, error) {
	var id int

	// Marshal complex types to JSON
	equipmentJSON, err := json.Marshal(v.Equipment)
	if err != nil {
		return 0, err
	}

	muscleJSON, err := json.Marshal(v.MuscleGroup)
	if err != nil {
		return 0, err
	}

	categoryJSON, err := json.Marshal(v.Category)
	if err != nil {
		return 0, err
	}

	err = WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		err := tx.QueryRowContext(
			ctx,
			CreateView,
			v.ID,
			v.Name,
			v.Description,
			v.Difficulty,
			v.Position,
			pq.Array(v.Alias),
			equipmentJSON,
			muscleJSON,
			categoryJSON,
			v.CreatedAt,
			v.UpdatedAt,
			v.DeletedAt,
			v.PurgeAt,
		).Scan(&id)
		if err != nil {
			return err
		}

		logr.Get().Info("Exercise view created!")
		return nil
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

const UpdateView = `
	UPDATE exercise_view
	SET name = $2,
		description = $3,
		difficulty = $4,
		position = $5,
		alias = $6,
		equipment = $7,
		attachment = $8,
		muscle = $9,
		category = $10,
		updated_at = $11,
		deleted_at = $12,
		purge_at = $13
	WHERE id = $1 AND deleted_at IS NULL
`

func (w *ViewWriter) Update(ctx context.Context, v view.View) error {
	// Marshal complex types to JSON
	equipmentJSON, err := json.Marshal(v.Equipment)
	if err != nil {
		return err
	}

	muscleJSON, err := json.Marshal(v.MuscleGroup)
	if err != nil {
		return err
	}

	categoryJSON, err := json.Marshal(v.Category)
	if err != nil {
		return err
	}

	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(
			ctx,
			UpdateView,
			v.ID,
			v.Name,
			v.Description,
			v.Difficulty,
			v.Position,
			pq.Array(v.Alias),
			equipmentJSON,
			muscleJSON,
			categoryJSON,
			v.UpdatedAt,
			v.DeletedAt,
			v.PurgeAt,
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

		logr.Get().Info("Exercise view updated!")
		return nil
	})
}

const DeleteView = `
	UPDATE exercise_view 
	SET deleted_at = CURRENT_TIMESTAMP,
		purge_at = CURRENT_TIMESTAMP + INTERVAL '45 days'
	WHERE id = $1 AND deleted_at IS NULL
`

func (w *ViewWriter) Delete(ctx context.Context, id int) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteView, id)
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

		logr.Get().Info("Exercise view soft deleted!")
		return nil
	})
}
