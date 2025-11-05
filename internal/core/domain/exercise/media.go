package exercise

import (
	"errors"
	"time"
)

var ErrInvalidName = errors.New("empty name")

type Media struct {
	URL       string
	Type      MediaType
	Order     int
	IsPrimary bool
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewMedia(url string, mediaType MediaType, order int, isPrimary bool) (Media, error) {
	return Media{
		URL:       url,
		Type:      mediaType,
		Order:     order,
		IsPrimary: isPrimary,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
