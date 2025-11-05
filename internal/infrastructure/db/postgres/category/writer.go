// Package category
package category

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

const CreateCategory = `INSERT INTO categories (name, created_at, updated_at) VALUES ($1, $2, $3)`

func (w *Writer) Add(ctx context.Context, category category.Category) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, CreateCategory, category.Name, category.CreatedAt, category.UpdatedAt)
		if err != nil {
			return err
		}

		logr.Get().Info("New category created!")
		return nil
	})
}

const UpdateCategory = `UPDATE categories
	SET name = $2, 
		updated_at = $3
	WHERE id = $1
`

func (w *Writer) Update(ctx context.Context, category category.Category) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, UpdateCategory, category.ID, category.Name, category.UpdatedAt)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return ports.ErrCategoryNotFound
		}

		logr.Get().Info("Category updated!")
		return nil
	})
}

const DeleteCategory = `DELETE FROM categories WHERE id = $1`

func (w *Writer) Delete(ctx context.Context, id int) error {
	return postgres.WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteCategory, id)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrCategoryNotFound
		}

		logr.Get().Info("Category deleted!")
		return nil
	})
}
