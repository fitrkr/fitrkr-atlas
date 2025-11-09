package equipments

import (
	"context"
	"fmt"
	"strings"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateAttachmentCommand struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Write ports.Write
	Read  ports.Read
}

type UpdateAttachmentResp struct{}

func (cmd *UpdateAttachmentCommand) Handle(ctx context.Context) (any, error) {
	existing, err := cmd.Read.Attachment.GetByID(ctx, cmd.ID)
	if err != nil {
		return UpdateAttachmentResp{}, fmt.Errorf("failed to get attachment: %w", err)
	}

	if cmd.Name != "" {
		name := strings.TrimSpace(strings.ToLower(cmd.Name))
		existing.Name = name
	}

	if cmd.Type != "" {
		attachmentType, err := equipment.NewAttachmentType(cmd.Type)
		if err != nil {
			return nil, fmt.Errorf("failed to create attachment type: %w", err)
		}
		existing.Type = attachmentType
	}

	existing.Touch()

	err = cmd.Write.Attachment.Update(ctx, *existing)
	if err != nil {
		return UpdateAttachmentResp{}, fmt.Errorf("failed to update attachment: %w", err)
	}

	return UpdateAttachmentResp{}, nil
}
