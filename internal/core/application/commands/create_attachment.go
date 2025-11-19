package commands

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
)

type CreateAttachmentCommand struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreateAttachmentResp struct{}

func (cmd *CreateAttachmentCommand) Handle(ctx context.Context) (any, error) {
	attachmentType, err := equipment.NewAttachmentType(cmd.Type)
	if err != nil {
		return nil, fmt.Errorf("failed to create new attachment type: %w", err)
	}

	attachment, err := equipment.NewAttachment(cmd.Name, attachmentType.ToString())
	if err != nil {
		return CreateAttachmentResp{}, fmt.Errorf("failed to create new attachment: %w", err)
	}

	err = write.Attachment.Add(ctx, attachment)
	if err != nil {
		return CreateAttachmentResp{}, fmt.Errorf("failed to add attachment: %w", err)
	}

	return CreateAttachmentResp{}, nil
}

func init() {
	register(&CreateAttachmentCommand{})
}
