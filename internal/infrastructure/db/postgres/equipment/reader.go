package equipment

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

const GetEquipmentByID = `SELECT id, name, description, created_at, updated_at FROM equipments WHERE id = $1`

func (r *Reader) GetByID(ctx context.Context, id int) (*equipment.Equipment, error) {
	var row equipment.Equipment

	err := r.db.QueryRowContext(ctx, GetEquipmentByID, id).Scan(
		&row.ID,
		&row.Name,
		&row.Description,
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

const GetEquipments = `SELECT id, name, description, created_at, updated_at FROM equipments`

func (r *Reader) GetAll(ctx context.Context) ([]equipment.Equipment, error) {
	rows, err := r.db.QueryContext(ctx, GetEquipments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipments []equipment.Equipment

	for rows.Next() {
		var eq equipment.Equipment
		err := rows.Scan(
			&eq.ID,
			&eq.Name,
			&eq.Description,
			&eq.CreatedAt,
			&eq.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		equipments = append(equipments, eq)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(equipments) == 0 {
		return nil, ports.ErrEquipmentNotFound
	}

	return equipments, nil
}
