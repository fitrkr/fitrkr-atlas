package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type AttachmentReader struct {
	db *sql.DB
}

func newAttachmentReader(db *sql.DB) *AttachmentReader {
	return &AttachmentReader{db: db}
}

const GetAttachmentByID = `SELECT id, name, type, created_at, updated_at FROM attachment WHERE id = $1`

func (r *AttachmentReader) GetByID(ctx context.Context, id int) (*equipment.Attachment, error) {
	var row equipment.Attachment

	err := r.db.QueryRowContext(ctx, GetAttachmentByID, id).Scan(
		&row.ID,
		&row.Name,
		&row.Type,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrAttachmentNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetAttachments = `SELECT id, name, type, created_at, updated_at FROM attachment`

func (r *AttachmentReader) GetAll(ctx context.Context) ([]*equipment.Attachment, error) {
	rows, err := r.db.QueryContext(ctx, GetAttachments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attachments []*equipment.Attachment

	for rows.Next() {
		var row equipment.Attachment
		err := rows.Scan(
			&row.ID,
			&row.Name,
			&row.Type,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		attachments = append(attachments, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return attachments, nil
}
