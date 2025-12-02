package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type ExerciseAttachmentReader struct {
	db *sql.DB
}

func newExerciseAttachmentReader(db *sql.DB) *ExerciseAttachmentReader {
	return &ExerciseAttachmentReader{db: db}
}

const GetExerciseAttachmentByID = `SELECT id, exercise_id, attachment_id, created_at, updated_at FROM exercise_attachment WHERE id = $1`

func (r *ExerciseAttachmentReader) GetByID(ctx context.Context, id int) (*exercise.ExerciseAttachment, error) {
	var row exercise.ExerciseAttachment

	err := r.db.QueryRowContext(ctx, GetExerciseAttachmentByID, id).Scan(
		&row.ID,
		&row.ExerciseID,
		&row.AttachmentID,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ports.ErrExerciseAttachmentNotFound
		}
		return nil, err
	}
	return &row, nil
}

const GetAttachmentsByExerciseID = `SELECT id, exercise_id, attachment_id, created_at, updated_at FROM exercise_attachment WHERE exercise_id = $1`

func (r *ExerciseAttachmentReader) GetByExerciseID(ctx context.Context, exerciseID int) ([]*exercise.ExerciseAttachment, error) {
	rows, err := r.db.QueryContext(ctx, GetAttachmentsByExerciseID, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attachments []*exercise.ExerciseAttachment

	for rows.Next() {
		var row exercise.ExerciseAttachment
		err := rows.Scan(
			&row.ID,
			&row.ExerciseID,
			&row.AttachmentID,
			&row.CreatedAt,
			&row.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		attachments = append(attachments, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return attachments, nil
}
