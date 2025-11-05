// Package equipment
package equipment

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
	"github.com/cheezecakee/fitrkr-atlas/internal/infrastructure/db/postgres"
)

type Writer struct {
	db *sql.DB
}

func NewWriter(db *sql.DB) *Writer {
	return &Writer{db: db}
}

const CreateEquipment = `INSERT INTO equipments (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4)`

func (w *Writer) Add(ctx context.Context, eq equipment.Equipment) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, CreateEquipment, eq.Name, eq.Description, eq.CreatedAt, eq.UpdatedAt)
		if err != nil {
			return err
		}

		logr.Get().Info("New equipment created!")
		return nil
	})
}

const UpdateEquipment = `UPDATE equipments
	SET name = $2, 
		description = $3,
		updated_at = $4
	WHERE id = $1
`

func (w *Writer) Update(ctx context.Context, eq equipment.Equipment) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, UpdateEquipment, eq.ID, eq.Name, eq.Description, eq.UpdatedAt)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return ports.ErrEquipmentNotFound
		}

		logr.Get().Info("Equipment updated!")
		return nil
	})
}

const DeleteEquipment = `DELETE FROM equipments WHERE id = $1`

func (w *Writer) Delete(ctx context.Context, id int) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteEquipment, id)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrEquipmentNotFound
		}

		logr.Get().Info("Equipment deleted!")
		return nil
	})
}
