package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllAttachmentsQuery struct {
	Read ports.Read
}

type GetAllAttachmentsResp struct {
	Attachments []equipment.Attachment
}

func (qry *GetAllAttachmentsQuery) Handle(ctx context.Context) (any, error) {
	attachments, err := qry.Read.Attachment.GetAll(ctx)
	if err != nil {
		return GetAllAttachmentsResp{}, fmt.Errorf("failed to get attachments: %w", err)
	}

	return GetAllAttachmentsResp{Attachments: attachments}, nil
}
