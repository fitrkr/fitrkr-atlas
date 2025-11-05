package attachment

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type Reader struct {
	db *sql.DB
}

func NewReader(db *sql.DB) *Reader {
	return &Reader{db: db}
}

const GetAttachmentByID = `SELECT id, equipment_id, name, created_at, updated_at FROM equipment_attachments WHERE id = $1`

func (r *Reader) GetByID(ctx context.Context, id int) (*equipment.Attachment, error) {
	var row equipment.Attachment

	err := r.db.QueryRowContext(ctx, GetAttachmentByID, id).Scan(
		&row.ID,
		&row.EquipmentID,
		&row.Name,
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

const GetAttachments = `SELECT id, equipment_id, name, created_at, updated_at FROM equipment_attachments`

func (r *Reader) GetAll(ctx context.Context) ([]equipment.Attachment, error) {
	rows, err := r.db.QueryContext(ctx, GetAttachments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attachments []equipment.Attachment

	for rows.Next() {
		var at equipment.Attachment
		err := rows.Scan(
			&at.ID,
			&at.EquipmentID,
			&at.Name,
			&at.CreatedAt,
			&at.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		attachments = append(attachments, at)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(attachments) == 0 {
		return nil, ports.ErrAttachmentNotFound
	}

	return attachments, nil
}

const GetAttachmentsByEquipmentID = `SELECT id, equipment_id,  name, created_at, updated_at FROM equipment_attachments WHERE equipment_id = $1`

func (r *Reader) GetByEquipmentID(ctx context.Context, equipmentID int) ([]equipment.Attachment, error) {
	rows, err := r.db.QueryContext(ctx, GetAttachmentsByEquipmentID, equipmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attachments []equipment.Attachment

	for rows.Next() {
		var at equipment.Attachment
		err := rows.Scan(
			&at.ID,
			&at.EquipmentID,
			&at.Name,
			&at.CreatedAt,
			&at.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		attachments = append(attachments, at)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(attachments) == 0 {
		return nil, ports.ErrAttachmentNotFound
	}

	return attachments, nil
}
