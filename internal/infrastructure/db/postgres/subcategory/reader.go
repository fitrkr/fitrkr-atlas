package subcategory

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

const GetSubcategoryByID = `SELECT id, category_id, name, created_at, updated_at FROM subcategories WHERE id = $1`

func (r *Reader) GetByID(ctx context.Context, id int) (*category.Subcategory, error) {
	var row category.Subcategory

	err := r.db.QueryRowContext(ctx, GetSubcategoryByID, id).Scan(
		&row.ID,
		&row.CategoryID,
		&row.Name,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrSubcategoryNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetSubcategories = `SELECT id, category_id, name, created_at, updated_at FROM subcategories`

func (r *Reader) GetAll(ctx context.Context) ([]category.Subcategory, error) {
	rows, err := r.db.QueryContext(ctx, GetSubcategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subcategories []category.Subcategory

	for rows.Next() {
		var subcategory category.Subcategory
		err := rows.Scan(
			&subcategory.ID,
			&subcategory.CategoryID,
			&subcategory.Name,
			&subcategory.CreatedAt,
			&subcategory.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		subcategories = append(subcategories, subcategory)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(subcategories) == 0 {
		return nil, ports.ErrSubcategoryNotFound
	}

	return subcategories, nil
}

const GetSubcategoriesByCategoryID = `SELECT id, category_id,  name, created_at, updated_at FROM subcategories WHERE category_id = $1`

func (r *Reader) GetByCategoryID(ctx context.Context, categoryID int) ([]category.Subcategory, error) {
	rows, err := r.db.QueryContext(ctx, GetSubcategoriesByCategoryID, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subcategories []category.Subcategory

	for rows.Next() {
		var subcategory category.Subcategory
		err := rows.Scan(
			&subcategory.ID,
			&subcategory.CategoryID,
			&subcategory.Name,
			&subcategory.CreatedAt,
			&subcategory.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		subcategories = append(subcategories, subcategory)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(subcategories) == 0 {
		return nil, ports.ErrSubcategoryNotFound
	}

	return subcategories, nil
}
