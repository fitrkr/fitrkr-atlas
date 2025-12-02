package queries

import (
	"context"
	"fmt"

	"github.com/fitrkr/atlas/internal/core/domain/exercise"
)

type GetExerciseAttachmentByIDQuery struct {
	ID int `json:"id"`
}

type GetExerciseAttachmentByIDResp struct {
	Attachment *exercise.ExerciseAttachment
}

func (qry *GetExerciseAttachmentByIDQuery) Handle(ctx context.Context) (any, error) {
	attachment, err := read.Exercise.Attachment.GetByID(ctx, qry.ID)
	if err != nil {
		return GetExerciseAttachmentByIDResp{}, fmt.Errorf("failed to get exercise attachment: %w", err)
	}

	return GetExerciseAttachmentByIDResp{Attachment: attachment}, nil
}

func init() {
	register(&GetExerciseAttachmentByIDQuery{})
}
