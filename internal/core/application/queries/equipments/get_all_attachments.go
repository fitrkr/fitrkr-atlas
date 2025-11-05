package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllAttachmentsQuery struct {
	Read ports.EquipmentAttachmentRead
}

type GetAllAttachmentsResp struct {
	Attachments []equipment.Attachment
}

func (qry *GetAllAttachmentsQuery) Handle(ctx context.Context) (any, error) {
	attachments, err := qry.Read.GetAll(ctx)
	if err != nil {
		logr.Get().Errorf("failed to get attachments: %v", err)
		return GetAllAttachmentsResp{}, fmt.Errorf("failed to get attachments: %w", err)
	}

	return GetAllAttachmentsResp{Attachments: attachments}, nil
}
