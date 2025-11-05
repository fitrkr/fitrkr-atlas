// Package musclegroup
package musclegroup

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

const CreateMuscleGroup = `INSERT INTO muscle_groups (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4)`

func (w *Writer) Add(ctx context.Context, muscle muscle.Group) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, CreateMuscleGroup, muscle.Name, muscle.Description, muscle.CreatedAt, muscle.UpdatedAt)
		if err != nil {
			logr.Get().Infof("Failed to create muscle: %v", err)
			return err
		}

		logr.Get().Info("New muscle group created!")
		return nil
	})
}

const UpdateMuscleGroup = `UPDATE muscle_groups
	SET name = $2, 
		description = $3,
		updated_at = $4
	WHERE id = $1
`

func (w *Writer) Update(ctx context.Context, muscle muscle.Group) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, UpdateMuscleGroup, muscle.ID, muscle.Name, muscle.Description, muscle.UpdatedAt)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return ports.ErrMuscleGroupNotFound
		}

		logr.Get().Info("Muscle group updated!")
		return nil
	})
}

const DeleteMuscleGroup = `DELETE FROM muscle_groups WHERE id = $1`

func (w *Writer) Delete(ctx context.Context, id int) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteMuscleGroup, id)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrMuscleGroupNotFound
		}

		logr.Get().Info("Muscle group deleted!")
		return nil
	})
}
