package musclegroup

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

const GetMuscleGroupByID = `SELECT id, name, description, created_at, updated_at FROM muscle_groups WHERE id = $1`

func (r *Reader) GetByID(ctx context.Context, id int) (*muscle.Group, error) {
	var row muscle.Group

	err := r.db.QueryRowContext(ctx, GetMuscleGroupByID, id).Scan(
		&row.ID,
		&row.Name,
		&row.Description,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrMuscleGroupNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetMuscleGroups = `SELECT id, name, description, created_at, updated_at FROM muscle_groups`

func (r *Reader) GetAll(ctx context.Context) ([]muscle.Group, error) {
	rows, err := r.db.QueryContext(ctx, GetMuscleGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var muscleGroups []muscle.Group

	for rows.Next() {
		var m muscle.Group
		err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.Description,
			&m.CreatedAt,
			&m.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		muscleGroups = append(muscleGroups, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(muscleGroups) == 0 {
		return nil, ports.ErrMuscleGroupNotFound
	}

	return muscleGroups, nil
}
