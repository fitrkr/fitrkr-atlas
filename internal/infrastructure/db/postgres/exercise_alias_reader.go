package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type ExerciseAliasReader struct {
	db *sql.DB
}

func newExerciseAliasReader(db *sql.DB) *ExerciseAliasReader {
	return &ExerciseAliasReader{db: db}
}

const GetExerciseAliasByID = `SELECT id, exercise_id, name, language_code, created_at, updated_at FROM exercise_alias WHERE id = $1`

func (r *ExerciseAliasReader) GetByID(ctx context.Context, id int) (*exercise.Alias, error) {
	var row exercise.Alias

	err := r.db.QueryRowContext(ctx, GetExerciseAliasByID, id).Scan(
		&row.ID,
		&row.ExerciseID,
		&row.Name,
		&row.LanguageCode,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrExerciseAliasNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetExerciseAliasByName = `SELECT id, exercise_id, name, language_code, created_at, updated_at FROM exercise_alias WHERE name = $1`

func (r *ExerciseAliasReader) GetByName(ctx context.Context, name string) (*exercise.Alias, error) {
	var row exercise.Alias

	err := r.db.QueryRowContext(ctx, GetExerciseAliasByName, name).Scan(
		&row.ID,
		&row.ExerciseID,
		&row.Name,
		&row.LanguageCode,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrExerciseAliasNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetAliasesByExerciseID = `SELECT id, exercise_id, name, language_code, created_at, updated_at FROM exercise_alias WHERE exercise_id = $1`

func (r *ExerciseAliasReader) GetByExerciseID(ctx context.Context, exerciseID int) ([]*exercise.Alias, error) {
	rows, err := r.db.QueryContext(ctx, GetAliasesByExerciseID, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aliases []*exercise.Alias

	for rows.Next() {
		var row exercise.Alias
		err := rows.Scan(
			&row.ID,
			&row.ExerciseID,
			&row.Name,
			&row.LanguageCode,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		aliases = append(aliases, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return aliases, nil
}
