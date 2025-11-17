package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/lib/pq"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/view"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type ViewReader struct {
	db *sql.DB
}

func newViewReader(db *sql.DB) *ViewReader {
	return &ViewReader{db: db}
}

const GetViewByID = `
	SELECT 
		id, name, description, difficulty, position,
		alias, equipment, attachment, muscle, category,
		created_at, updated_at, deleted_at, purge_at
	FROM exercise_view
	WHERE id = $1 AND deleted_at IS NULL
`

func (r *ViewReader) GetByID(ctx context.Context, id int) (*view.View, error) {
	var (
		v             view.View
		equipmentJSON []byte
		muscleJSON    []byte
		categoryJSON  []byte
	)

	err := r.db.QueryRowContext(ctx, GetViewByID, id).Scan(
		&v.ID,
		&v.Name,
		&v.Description,
		&v.Difficulty,
		&v.Position,
		pq.Array(&v.Alias),
		&equipmentJSON,
		&muscleJSON,
		&categoryJSON,
		&v.CreatedAt,
		&v.UpdatedAt,
		&v.DeletedAt,
		&v.PurgeAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrExerciseNotFound
		}
		return nil, err
	}

	// Unmarshal JSONB fields
	if err := json.Unmarshal(equipmentJSON, &v.Equipment); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(muscleJSON, &v.MuscleGroup); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(categoryJSON, &v.Category); err != nil {
		return nil, err
	}

	return &v, nil
}

const GetAllViews = `
	SELECT 
		id, name, description, difficulty, position,
		alias, equipment, attachment, muscle, category,
		created_at, updated_at, deleted_at, purge_at
	FROM exercise_view
	WHERE deleted_at IS NULL
	ORDER BY name
`

func (r *ViewReader) GetAll(ctx context.Context) ([]*view.View, error) {
	rows, err := r.db.QueryContext(ctx, GetAllViews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []*view.View

	for rows.Next() {
		var v view.View
		var equipmentJSON []byte
		var muscleJSON []byte
		var categoryJSON []byte

		err := rows.Scan(
			&v.ID,
			&v.Name,
			&v.Description,
			&v.Difficulty,
			&v.Position,
			pq.Array(&v.Alias),
			&equipmentJSON,
			&muscleJSON,
			&categoryJSON,
			&v.CreatedAt,
			&v.UpdatedAt,
			&v.DeletedAt,
			&v.PurgeAt,
		)
		if err != nil {
			return nil, err
		}

		// Unmarshal JSONB fields
		if err := json.Unmarshal(equipmentJSON, &v.Equipment); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(muscleJSON, &v.MuscleGroup); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(categoryJSON, &v.Category); err != nil {
			return nil, err
		}

		views = append(views, &v)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return views, nil
}
