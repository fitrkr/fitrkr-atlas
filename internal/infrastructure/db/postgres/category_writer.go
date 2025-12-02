package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"

	"github.com/fitrkr/atlas/internal/core/domain/category"
	"github.com/fitrkr/atlas/internal/core/ports"
)

type CategoryWriter struct {
	db *sql.DB
}

func newCategoryWriter(db *sql.DB) *CategoryWriter {
	return &CategoryWriter{db: db}
}

const CreateCategory = `INSERT INTO category (name, type, created_at, updated_at) VALUES ($1, $2, $3, $4)`

func (w *CategoryWriter) Add(ctx context.Context, category category.Category) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, CreateCategory, category.Name, category.Type, category.CreatedAt, category.UpdatedAt)
		if err != nil {
			return err
		}

		logr.Get().Info("Category created!")
		return nil
	})
}

const UpdateCategory = `UPDATE category
	SET name = $2, 
		type = $3,
		updated_at = $4
	WHERE id = $1
`

func (w *CategoryWriter) Update(ctx context.Context, category category.Category) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, UpdateCategory, category.ID, category.Name, category.Type, category.UpdatedAt)
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

const DeleteCategory = `DELETE FROM category WHERE id = $1`

func (w *CategoryWriter) Delete(ctx context.Context, id int) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
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
