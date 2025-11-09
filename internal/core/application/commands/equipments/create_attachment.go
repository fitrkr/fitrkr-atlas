package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateAttachmentCommand struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Write ports.Write
	Read  ports.Read
}

type CreateAttachmentResp struct{}

func (cmd *CreateAttachmentCommand) Handle(ctx context.Context) (any, error) {
	attachmentType, err := equipment.NewAttachmentType(cmd.Type)
	if err != nil {
		return nil, fmt.Errorf("failed to create new attachment type: %w", err)
	}

	attachment, err := equipment.NewAttachment(cmd.Name, attachmentType)
	if err != nil {
		return CreateAttachmentResp{}, fmt.Errorf("failed to create new attachment: %w", err)
	}

	err = cmd.Write.Attachment.Add(ctx, attachment)
	if err != nil {
		return CreateAttachmentResp{}, fmt.Errorf("failed to add attachment: %w", err)
	}

	return CreateAttachmentResp{}, nil
}
