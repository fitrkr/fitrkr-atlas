package exercise

import (
	"errors"
	"time"
)

var ErrEmptyAttachmentID = errors.New("empty equipment id")

type ExerciseAttachment struct {
	ID           *int      `json:"id"`
	ExerciseID   int       `json:"exercise_id"`
	AttachmentID *int      `json:"attachment_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewExerciseAttachment(exerciseID int, attachmentID *int) (*ExerciseAttachment, error) {
	if exerciseID < 0 {
		return nil, ErrEmptyExericiseID
	}
	if attachmentID == nil {
		return &ExerciseAttachment{}, nil // Not an error
	}
	return &ExerciseAttachment{
		ExerciseID:   exerciseID,
		AttachmentID: attachmentID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}
