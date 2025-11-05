// Package attachment
package attachment

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

const CreateAttachment = `INSERT INTO equipment_attachments (equipment_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)`

func (w *Writer) Add(ctx context.Context, at equipment.Attachment) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, CreateAttachment, at.EquipmentID, at.Name, at.CreatedAt, at.UpdatedAt)
		if err != nil {
			return err
		}

		logr.Get().Info("New attachment created!")
		return nil
	})
}

const UpdateAttachment = `UPDATE equipment_attachments
	SET equipment_id = $2, 
		name = $3,
		updated_at = $4
	WHERE id = $1
`

func (w *Writer) Update(ctx context.Context, at equipment.Attachment) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, UpdateAttachment, at.ID, at.EquipmentID, at.Name, at.UpdatedAt)
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

		logr.Get().Info("Attachment updated!")
		return nil
	})
}

const DeleteAttachment = `DELETE FROM equipment_attachments WHERE id = $1`

func (w *Writer) Delete(ctx context.Context, id int) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteAttachment, id)
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

		logr.Get().Info("Attachment deleted!")
		return nil
	})
}
