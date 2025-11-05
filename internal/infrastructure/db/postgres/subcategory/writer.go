// Package subcategory
package subcategory

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
	"github.com/cheezecakee/fitrkr-atlas/internal/infrastructure/db/postgres"
)

type Writer struct {
	db *sql.DB
}

func NewWriter(db *sql.DB) *Writer {
	return &Writer{db: db}
}

const CreateSubcategory = `INSERT INTO subcategories (category_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)`

func (w *Writer) Add(ctx context.Context, subcategory category.Subcategory) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, CreateSubcategory, subcategory.CategoryID, subcategory.Name, subcategory.CreatedAt, subcategory.UpdatedAt)
		if err != nil {
			return err
		}

		logr.Get().Info("New subcategory created!")
		return nil
	})
}

const UpdateSubcategory = `UPDATE subcategories
	SET category_id = $2, 
		name = $3,
		updated_at = $4
	WHERE id = $1
`

func (w *Writer) Update(ctx context.Context, subcategory category.Subcategory) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, UpdateSubcategory, subcategory.ID, subcategory.CategoryID, subcategory.Name, subcategory.UpdatedAt)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return ports.ErrSubcategoryNotFound
		}

		logr.Get().Info("Subcategory updated!")
		return nil
	})
}

const DeleteSubcategory = `DELETE FROM subcategories WHERE id = $1`

func (w *Writer) Delete(ctx context.Context, id int) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteSubcategory, id)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrSubcategoryNotFound
		}

		logr.Get().Info("Subcategory deleted!")
		return nil
	})
}
