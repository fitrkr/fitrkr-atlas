package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type ExerciseMuscleReader struct {
	db *sql.DB
}

func newExerciseMuscleReader(db *sql.DB) *ExerciseMuscleReader {
	return &ExerciseMuscleReader{db: db}
}

const GetExerciseMuscleByID = `SELECT id, exercise_id, muscle_id, activation, created_at, updated_at FROM exercise_muscle WHERE id = $1`

func (r *ExerciseMuscleReader) GetByID(ctx context.Context, id int) (*exercise.ExerciseMuscle, error) {
	var row exercise.ExerciseMuscle

	err := r.db.QueryRowContext(ctx, GetExerciseMuscleByID, id).Scan(
		&row.ID,
		&row.ExerciseID,
		&row.MuscleID,
		&row.Activation,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrExerciseMuscleNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetMusclesByExerciseID = `SELECT id, exercise_id, muscle_id, activation, created_at, updated_at FROM exercise_muscle WHERE exercise_id = $1`

func (r *ExerciseMuscleReader) GetByExerciseID(ctx context.Context, exerciseID int) ([]*exercise.ExerciseMuscle, error) {
	rows, err := r.db.QueryContext(ctx, GetMusclesByExerciseID, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var muscles []*exercise.ExerciseMuscle

	for rows.Next() {
		var row exercise.ExerciseMuscle
		err := rows.Scan(
			&row.ID,
			&row.ExerciseID,
			&row.MuscleID,
			&row.Activation,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		muscles = append(muscles, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return muscles, nil
}
