package category

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type Reader struct {
	db *sql.DB
}

func NewReader(db *sql.DB) *Reader {
	return &Reader{db: db}
}

const GetCategoryByID = `SELECT id, name, created_at, updated_at FROM categories WHERE id = $1`

func (r *Reader) GetByID(ctx context.Context, id int) (*category.Category, error) {
	var row category.Category

	err := r.db.QueryRowContext(ctx, GetCategoryByID, id).Scan(
		&row.ID,
		&row.Name,
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

const GetCategories = `SELECT id, name, created_at, updated_at FROM categories`

func (r *Reader) GetAll(ctx context.Context) ([]category.Category, error) {
	rows, err := r.db.QueryContext(ctx, GetCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []category.Category

	for rows.Next() {
		var category category.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, ports.ErrCategoryNotFound
	}

	return categories, nil
}
