package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/equipment"
)

type GetAllAttachmentsQuery struct{}

type GetAllAttachmentsResp struct {
	Attachments []*equipment.Attachment
}

func (qry *GetAllAttachmentsQuery) Handle(ctx context.Context) (any, error) {
	attachments, err := read.Attachment.GetAll(ctx)
	if err != nil {
		return GetAllAttachmentsResp{}, fmt.Errorf("failed to get attachments: %w", err)
	}

	return GetAllAttachmentsResp{Attachments: attachments}, nil
}

func init() {
	register(&GetAllAttachmentsQuery{})
}
