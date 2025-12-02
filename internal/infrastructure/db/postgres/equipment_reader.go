package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type EquipmentReader struct {
	db *sql.DB
}

func newEquipmentReader(db *sql.DB) *EquipmentReader {
	return &EquipmentReader{db: db}
}

const GetEquipmentByID = `SELECT id, name, description, type, created_at, updated_at FROM equipment WHERE id = $1`

func (r *EquipmentReader) GetByID(ctx context.Context, id int) (*equipment.Equipment, error) {
	var row equipment.Equipment

	err := r.db.QueryRowContext(ctx, GetEquipmentByID, id).Scan(
		&row.ID,
		&row.Name,
		&row.Description,
		&row.Type,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrEquipmentNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetEquipments = `SELECT id, name, description, type, created_at, updated_at FROM equipment`

func (r *EquipmentReader) GetAll(ctx context.Context) ([]*equipment.Equipment, error) {
	rows, err := r.db.QueryContext(ctx, GetEquipments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipments []*equipment.Equipment

	for rows.Next() {
		var row equipment.Equipment
		err := rows.Scan(
			&row.ID,
			&row.Name,
			&row.Description,
			&row.Type,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		equipments = append(equipments, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return equipments, nil
}
