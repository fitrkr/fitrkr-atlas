package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CategoryReader struct {
	db *sql.DB
}

func newCategoryReader(db *sql.DB) *CategoryReader {
	return &CategoryReader{db: db}
}

const GetCategoryByID = `SELECT id, name, type, created_at, updated_at FROM category WHERE id = $1`

func (r *CategoryReader) GetByID(ctx context.Context, id int) (*category.Category, error) {
	var row category.Category

	err := r.db.QueryRowContext(ctx, GetCategoryByID, id).Scan(
		&row.ID,
		&row.Name,
		&row.Type,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrCategoryNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetCategories = `SELECT id, name, type, created_at, updated_at FROM category`

func (r *CategoryReader) GetAll(ctx context.Context) ([]*category.Category, error) {
	rows, err := r.db.QueryContext(ctx, GetCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*category.Category

	for rows.Next() {
		var row category.Category
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
		categories = append(categories, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

const GetCategoriesByType = `SELECT id, name, type, created_at, updated_at FROM category WHERE type = $1`

func (r *CategoryReader) GetByType(ctx context.Context, categoryType string) ([]*category.Category, error) {
	rows, err := r.db.QueryContext(ctx, GetCategoriesByType, categoryType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*category.Category

	for rows.Next() {
		var row category.Category
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
		categories = append(categories, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
