package muscle

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type Reader struct {
	db *sql.DB
}

func NewReader(db *sql.DB) *Reader {
	return &Reader{db: db}
}

const GetMuscleByID = `SELECT id, muscle_group_id, name, created_at, updated_at FROM muscles WHERE id = $1`

func (r *Reader) GetByID(ctx context.Context, id int) (*muscle.Muscle, error) {
	var row muscle.Muscle

	err := r.db.QueryRowContext(ctx, GetMuscleByID, id).Scan(
		&row.ID,
		&row.MuscleGroupID,
		&row.Name,
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

const GetMuscles = `SELECT id, muscle_group_id, name, created_at, updated_at FROM muscles`

func (r *Reader) GetAll(ctx context.Context) ([]muscle.Muscle, error) {
	rows, err := r.db.QueryContext(ctx, GetMuscles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var muscles []muscle.Muscle

	for rows.Next() {
		var m muscle.Muscle
		err := rows.Scan(
			&m.ID,
			&m.MuscleGroupID,
			&m.Name,
			&m.CreatedAt,
			&m.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		muscles = append(muscles, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(muscles) == 0 {
		return nil, ports.ErrMuscleNotFound
	}

	return muscles, nil
}

const GetMusclesByGroupID = `SELECT id, muscle_group_id,  name, created_at, updated_at FROM muscles WHERE muscle_group_id = $1`

func (r *Reader) GetByMuscleGroupID(ctx context.Context, muscleGroupID int) ([]muscle.Muscle, error) {
	rows, err := r.db.QueryContext(ctx, GetMusclesByGroupID, muscleGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var muscles []muscle.Muscle

	for rows.Next() {
		var m muscle.Muscle
		err := rows.Scan(
			&m.ID,
			&m.MuscleGroupID,
			&m.Name,
			&m.CreatedAt,
			&m.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		muscles = append(muscles, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(muscles) == 0 {
		return nil, ports.ErrMuscleNotFound
	}

	return muscles, nil
}
