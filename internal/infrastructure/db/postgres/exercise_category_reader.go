package postgres

import (
	"context"
	"database/sql"

	"github.com/fitrkr/atlas/internal/core/domain/exercise"
	"github.com/fitrkr/atlas/internal/core/ports"
)

type ExerciseCategoryReader struct {
	db *sql.DB
}

func newExerciseCategoryReader(db *sql.DB) *ExerciseCategoryReader {
	return &ExerciseCategoryReader{db: db}
}

const GetExerciseCategoryByID = `SELECT id, exercise_id, category_id, created_at, updated_at FROM exercise_category WHERE id = $1`

func (r *ExerciseCategoryReader) GetByID(ctx context.Context, id int) (*exercise.ExerciseCategory, error) {
	var row exercise.ExerciseCategory

	err := r.db.QueryRowContext(ctx, GetExerciseCategoryByID, id).Scan(
		&row.ID,
		&row.ExerciseID,
		&row.CategoryID,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrExerciseCategoryNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetCategoriesByExerciseID = `SELECT id, exercise_id, category_id, created_at, updated_at FROM exercise_category WHERE exercise_id = $1`

func (r *ExerciseCategoryReader) GetByExerciseID(ctx context.Context, exerciseID int) ([]*exercise.ExerciseCategory, error) {
	rows, err := r.db.QueryContext(ctx, GetCategoriesByExerciseID, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*exercise.ExerciseCategory

	for rows.Next() {
		var row exercise.ExerciseCategory
		err := rows.Scan(
			&row.ID,
			&row.ExerciseID,
			&row.CategoryID,
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
