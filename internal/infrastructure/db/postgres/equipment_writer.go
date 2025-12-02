package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"

	"github.com/fitrkr/atlas/internal/core/domain/equipment"
	"github.com/fitrkr/atlas/internal/core/ports"
)

type EquipmentWriter struct {
	db *sql.DB
}

func newEquipmentWriter(db *sql.DB) *EquipmentWriter {
	return &EquipmentWriter{db: db}
}

const CreateEquipment = `
	INSERT INTO equipment (name, description, type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) 
	RETURNING id
`

func (w *EquipmentWriter) Add(ctx context.Context, eq equipment.Equipment) (int, error) {
	var id int

	err := WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		err := tx.QueryRowContext(
			ctx,
			CreateEquipment,
			eq.Name,
			eq.Description,
			eq.Type,
			eq.CreatedAt,
			eq.UpdatedAt,
		).Scan(&id)
		if err != nil {
			return err
		}

		logr.Get().Info("Equipment created!")
		return nil
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

const UpdateEquipment = `
	UPDATE equipment
	SET name = $2, 
		description = $3,
		type = $4,
		updated_at = $5
	WHERE id = $1
`

func (w *EquipmentWriter) Update(ctx context.Context, eq equipment.Equipment) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, UpdateEquipment, eq.ID, eq.Name, eq.Description, eq.Type, eq.UpdatedAt)
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

const DeleteEquipment = `DELETE FROM equipment WHERE id = $1`

func (w *EquipmentWriter) Delete(ctx context.Context, id int) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
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
