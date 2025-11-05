package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAttachmentByIDQuery struct {
	ID   int `json:"id"`
	Read ports.EquipmentAttachmentRead
}

type GetAttachmentByIDResp struct {
	Attachment *equipment.Attachment
}

func (qry *GetAttachmentByIDQuery) Handle(ctx context.Context) (any, error) {
	attachment, err := qry.Read.GetByID(ctx, qry.ID)
	if err != nil {
		logr.Get().Errorf("failed to get attachment: %v", err)
		return GetAttachmentByIDResp{}, fmt.Errorf("failed to get attachment: %w", err)
	}

	return GetAttachmentByIDResp{Attachment: attachment}, nil
}
