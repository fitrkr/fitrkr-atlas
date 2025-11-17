package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type AttachmentWriter struct {
	db *sql.DB
}

func newAttachmentWriter(db *sql.DB) *AttachmentWriter {
	return &AttachmentWriter{db: db}
}

const CreateAttachment = `INSERT INTO attachment (name, type, created_at, updated_at) VALUES ($1, $2, $3, $4)`

func (w *AttachmentWriter) Add(ctx context.Context, at equipment.Attachment) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, CreateAttachment, at.Name, at.Type, at.CreatedAt, at.UpdatedAt)
		if err != nil {
			return err
		}

		logr.Get().Info("Attachment created!")
		return nil
	})
}

const UpdateAttachment = `UPDATE attachment
	SET name = $2, 
		type = $3,
		updated_at = $4
	WHERE id = $1
`

func (w *AttachmentWriter) Update(ctx context.Context, at equipment.Attachment) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, UpdateAttachment, at.ID, at.Name, at.Type, at.UpdatedAt)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return ports.ErrAttachmentNotFound
		}

		logr.Get().Info("Attachment updated!")
		return nil
	})
}

const DeleteAttachment = `DELETE FROM attachment WHERE id = $1`

func (w *AttachmentWriter) Delete(ctx context.Context, id int) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteAttachment, id)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrAttachmentNotFound
		}

		logr.Get().Info("Attachment deleted!")
		return nil
	})
}
