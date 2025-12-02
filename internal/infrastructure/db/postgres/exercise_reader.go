package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type ExerciseReader struct {
	db *sql.DB
}

func newExerciseReader(db *sql.DB) *ExerciseReader {
	return &ExerciseReader{db: db}
}

const GetExerciseByID = `SELECT id, name, description, difficulty, position, equipment_id, created_at, updated_at, deleted_at, purge_at FROM exercise WHERE id = $1`

func (r *ExerciseReader) GetByID(ctx context.Context, id int) (*exercise.Exercise, error) {
	var row exercise.Exercise

	err := r.db.QueryRowContext(ctx, GetExerciseByID, id).Scan(
		&row.ID,
		&row.Name,
		&row.Description,
		&row.Difficulty,
		&row.Position,
		&row.EquipmentID,
		&row.CreatedAt,
		&row.UpdatedAt,
		&row.DeletedAt,
		&row.PurgeAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrExerciseNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetExerciseByName = `SELECT id, name, description, difficulty, position, equipment_id, created_at, updated_at, deleted_at, purge_at FROM exercise WHERE name = $1`

func (r *ExerciseReader) GetByName(ctx context.Context, name string) (*exercise.Exercise, error) {
	var row exercise.Exercise

	err := r.db.QueryRowContext(ctx, GetExerciseByName, name).Scan(
		&row.ID,
		&row.Name,
		&row.Description,
		&row.Difficulty,
		&row.Position,
		&row.EquipmentID,
		&row.CreatedAt,
		&row.UpdatedAt,
		&row.DeletedAt,
		&row.PurgeAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrExerciseNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetExercises = `
	SELECT id, name, description, difficulty, position, equipment_id, created_at, updated_at, deleted_at, purge_at 	  FROM exercise 
	WHERE deleted_at IS NULL
	ORDER BY name
	`

func (r *ExerciseReader) GetAll(ctx context.Context) ([]*exercise.Exercise, error) {
	rows, err := r.db.QueryContext(ctx, GetExercises)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []*exercise.Exercise

	for rows.Next() {
		var row exercise.Exercise
		err := rows.Scan(
			&row.ID,
			&row.Name,
			&row.Description,
			&row.Difficulty,
			&row.Position,
			&row.EquipmentID,
			&row.CreatedAt,
			&row.UpdatedAt,
			&row.DeletedAt,
			&row.PurgeAt,
		)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return exercises, nil
}
