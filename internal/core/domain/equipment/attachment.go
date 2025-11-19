package equipment

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrEmptyAttachment       = errors.New("empty attachment")
	ErrInvalidEquipmentID    = errors.New("invalid equipment id")
	ErrInvalidAttachmentType = errors.New("invalid attachment type")
)

type Attachment struct {
	ID        *int
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAttachment(name, attachmentType string) (Attachment, error) {
	if name == "" {
		return Attachment{}, ErrEmptyAttachment
	}
	return Attachment{
		Name: strings.TrimSpace(strings.ToLower(name)),
		Type: attachmentType,
	}, nil
}

func (a *Attachment) Touch() {
	a.UpdatedAt = time.Now()
}
