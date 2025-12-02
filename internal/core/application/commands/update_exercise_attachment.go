package commands

import (
	"context"
	"fmt"

	"github.com/fitrkr/atlas/internal/core/domain/exercise"
)

type UpdateExerciseAttachmentCommand struct {
	Remove     []int `json:"remove,omitempty"`
	Add        []int `json:"add,omitempty"` // attachment_ids
	ExerciseID int   `json:"exercise_id"`
}

type UpdateExerciseAttachmentResp struct{}

func (cmd *UpdateExerciseAttachmentCommand) Handle(ctx context.Context) (any, error) {
	if len(cmd.Remove) != 0 {
		err := write.Exercise.Attachment.Delete(ctx, cmd.Remove)
		if err != nil {
			return nil, fmt.Errorf("failed to delete exercise attachments: %w", err)
		}
	}

	if len(cmd.Add) != 0 {
		var attachments []exercise.ExerciseAttachment
		for _, attachmentID := range cmd.Add {
			attachment, err := exercise.NewExerciseAttachment(cmd.ExerciseID, &attachmentID)
			if err != nil {
				return nil, fmt.Errorf("failed to create exercise attachment: %w", err)
			}

			attachments = append(attachments, *attachment)
		}

		err := write.Exercise.Attachment.Add(ctx, attachments)
		if err != nil {
			return nil, fmt.Errorf("failed to insert exercise attachments: %w", err)
		}
	}

	return UpdateExerciseAttachmentResp{}, nil
}

func init() {
	register(&UpdateExerciseAttachmentCommand{})
}
