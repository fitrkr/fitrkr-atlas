// Package muscle
package muscle

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
	"github.com/cheezecakee/fitrkr-atlas/internal/infrastructure/db/postgres"
)

type Writer struct {
	db *sql.DB
}

func NewWriter(db *sql.DB) *Writer {
	return &Writer{db: db}
}

const CreateMuscle = `INSERT INTO muscles (muscle_group_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)`

func (w *Writer) Add(ctx context.Context, muscle muscle.Muscle) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, CreateMuscle, muscle.MuscleGroupID, muscle.Name, muscle.CreatedAt, muscle.UpdatedAt)
		if err != nil {
			return err
		}

		logr.Get().Info("New muscle created!")
		return nil
	})
}

const UpdateMuscle = `UPDATE muscles
	SET muscle_group_id = $2, 
		name = $3,
		updated_at = $4
	WHERE id = $1
`

func (w *Writer) Update(ctx context.Context, muscle muscle.Muscle) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, UpdateMuscle, muscle.ID, muscle.MuscleGroupID, muscle.Name, muscle.UpdatedAt)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return ports.ErrMuscleNotFound
		}

		logr.Get().Info("Muscle updated!")
		return nil
	})
}

const DeleteMuscle = `DELETE FROM muscles WHERE id = $1`

func (w *Writer) Delete(ctx context.Context, id int) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteMuscle, id)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrMuscleNotFound
		}

		logr.Get().Info("Muscle deleted!")
		return nil
	})
}
