package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type MuscleWriter struct {
	db *sql.DB
}

func newMuscleWriter(db *sql.DB) *MuscleWriter {
	return &MuscleWriter{db: db}
}

const CreateMuscle = `INSERT INTO muscle (name, muscle_group, created_at, updated_at) VALUES ($1, $2, $3, $4)`

func (w *MuscleWriter) Add(ctx context.Context, muscle muscle.Muscle) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, CreateMuscle, muscle.Name, muscle.Group, muscle.CreatedAt, muscle.UpdatedAt)
		if err != nil {
			return err
		}

		logr.Get().Info("Muscle created!")
		return nil
	})
}

const UpdateMuscle = `UPDATE muscle
	SET name = $2, 
		muscle_group = $3,
		updated_at = $4
	WHERE id = $1
`

func (w *MuscleWriter) Update(ctx context.Context, muscle muscle.Muscle) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, UpdateMuscle, muscle.ID, muscle.Name, muscle.Group, muscle.UpdatedAt)
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

const DeleteMuscle = `DELETE FROM muscle WHERE id = $1`

func (w *MuscleWriter) Delete(ctx context.Context, id int) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
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
