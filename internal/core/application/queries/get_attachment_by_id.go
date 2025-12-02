package queries

import (
	"context"
	"fmt"

	"github.com/fitrkr/atlas/internal/core/domain/equipment"
)

type GetAttachmentByIDQuery struct {
	ID int `json:"id"`
}

type GetAttachmentByIDResp struct {
	Attachment *equipment.Attachment
}

func (qry *GetAttachmentByIDQuery) Handle(ctx context.Context) (any, error) {
	attachment, err := read.Attachment.GetByID(ctx, qry.ID)
	if err != nil {
		return GetAttachmentByIDResp{}, fmt.Errorf("failed to get attachment: %w", err)
	}

	return GetAttachmentByIDResp{Attachment: attachment}, nil
}

func init() {
	register(&GetAttachmentByIDQuery{})
}
