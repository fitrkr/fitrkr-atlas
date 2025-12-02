package postgres

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/logr"
	"github.com/lib/pq"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type ExerciseAttachmentWriter struct {
	db *sql.DB
}

func newExerciseAttachmentWriter(db *sql.DB) *ExerciseAttachmentWriter {
	return &ExerciseAttachmentWriter{db: db}
}

const CreateExerciseAttachments = `INSERT INTO exercise_attachment (exercise_id, attachment_id, created_at, updated_at) VALUES ($1, $2, $3, $4)`

func (w *ExerciseAttachmentWriter) Add(ctx context.Context, attachments []exercise.ExerciseAttachment) error {
	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		stmt, err := tx.PrepareContext(ctx, CreateExerciseAttachments)
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, attachment := range attachments {
			_, err := stmt.ExecContext(ctx,
				attachment.ExerciseID,
				attachment.AttachmentID,
				attachment.CreatedAt,
				attachment.UpdatedAt,
			)
			if err != nil {
				return err
			}
		}

		logr.Get().Infof("Created %d attachments!", len(attachments))
		return nil
	})
}

const DeleteExerciseAttachments = `DELETE FROM exercise_attachment WHERE id = ANY($1)`

func (w *ExerciseAttachmentWriter) Delete(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	return WithTransaction(ctx, w.db, func(tx *sql.Tx) error {
		result, err := tx.ExecContext(ctx, DeleteExerciseAttachments, pq.Array(ids))
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ports.ErrExerciseAttachmentNotFound
		}

		logr.Get().Infof("Deleted %d attachments!", rowsAffected)
		return nil
	})
}
