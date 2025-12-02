package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type EquipmentAttachmentWriter struct {
	db *sql.DB
}

func newEquipmentAttachmentWriter(db *sql.DB) *EquipmentAttachmentWriter {
	return &EquipmentAttachmentWriter{db: db}
}

const CreateEquipmentAttachment = `INSERT INTO equipment_attachment (equipment_id, attachment_id, created_at, updated_at) VALUES ($1, $2, $3, $4)`

func (w *EquipmentAttachmentWriter) Add(ctx context.Context, at equipment.EquipmentAttachment) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, CreateEquipmentAttachment, at.EquipmentID, at.AttachmentID, at.CreatedAt, at.UpdatedAt)
		if err != nil {
			return err
		}

		logr.Get().Info("Equipment attachment created!")
		return nil
	})
}

const DeleteEquipmentAttachment = `DELETE FROM equipment_attachment WHERE id = $1`

func (w *EquipmentAttachmentWriter) Delete(ctx context.Context, id int) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteEquipmentAttachment, id)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrEquipmentAttachmentNotFound
		}

		logr.Get().Info("Equipment attachment deleted!")
		return nil
	})
}
