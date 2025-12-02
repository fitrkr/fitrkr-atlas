package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type MuscleReader struct {
	db *sql.DB
}

func newMuscleReader(db *sql.DB) *MuscleReader {
	return &MuscleReader{db: db}
}

const GetMuscleByID = `SELECT id, name, muscle_group, created_at, updated_at FROM muscle WHERE id = $1`

func (r *MuscleReader) GetByID(ctx context.Context, id int) (*muscle.Muscle, error) {
	var row muscle.Muscle

	err := r.db.QueryRowContext(ctx, GetMuscleByID, id).Scan(
		&row.ID,
		&row.Name,
		&row.Group,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrMuscleNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetMuscles = `SELECT id, name, muscle_group, created_at, updated_at FROM muscle`

func (r *MuscleReader) GetAll(ctx context.Context) ([]*muscle.Muscle, error) {
	rows, err := r.db.QueryContext(ctx, GetMuscles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var muscles []*muscle.Muscle

	for rows.Next() {
		var row muscle.Muscle
		err := rows.Scan(
			&row.ID,
			&row.Name,
			&row.Group,
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

const GetMusclesByGroup = `SELECT id,  name, muscle_group, created_at, updated_at FROM muscle WHERE muscle_group = $1`

func (r *MuscleReader) GetByGroupType(ctx context.Context, muscleType string) ([]*muscle.Muscle, error) {
	rows, err := r.db.QueryContext(ctx, GetMusclesByGroup, muscleType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var muscles []*muscle.Muscle

	for rows.Next() {
		var row muscle.Muscle
		err := rows.Scan(
			&row.ID,
			&row.Name,
			&row.Group,
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
