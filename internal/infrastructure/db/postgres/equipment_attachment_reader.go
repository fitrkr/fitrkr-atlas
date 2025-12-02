package postgres

import (
	"context"
	"database/sql"

	"github.com/fitrkr/atlas/internal/core/domain/equipment"
	"github.com/fitrkr/atlas/internal/core/ports"
)

type EquipmentAttachmentReader struct {
	db *sql.DB
}

func newEquipmentAttachmentReader(db *sql.DB) *EquipmentAttachmentReader {
	return &EquipmentAttachmentReader{db: db}
}

const GetEquipmentAttachmentByID = `SELECT id, equipment_id, attachment_id, created_at, updated_at FROM equipment_attachment WHERE id = $1`

func (r *EquipmentAttachmentReader) GetByID(ctx context.Context, id int) (*equipment.EquipmentAttachment, error) {
	var row equipment.EquipmentAttachment

	err := r.db.QueryRowContext(ctx, GetEquipmentAttachmentByID, id).Scan(
		&row.ID,
		&row.EquipmentID,
		&row.AttachmentID,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrEquipmentAttachmentNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetEquipmentAttachmentsByEquipmentID = `SELECT id, equipment_id,  attachment_id, created_at, updated_at FROM equipment_attachment WHERE equipment_id = $1`

func (r *EquipmentAttachmentReader) GetByEquipmentID(ctx context.Context, equipmentID int) ([]*equipment.EquipmentAttachment, error) {
	rows, err := r.db.QueryContext(ctx, GetEquipmentAttachmentsByEquipmentID, equipmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attachments []*equipment.EquipmentAttachment

	for rows.Next() {
		var row equipment.EquipmentAttachment
		err := rows.Scan(
			&row.ID,
			&row.EquipmentID,
			&row.AttachmentID,
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
