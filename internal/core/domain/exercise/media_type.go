package exercise

import (
	"errors"
	"strings"
)

type MediaType string

const (
	MediaImage MediaType = "image"
	MediaGif   MediaType = "gif"
	MediaVideo MediaType = "video"
)

var ErrInvalidMediaType = errors.New("invalid media type")

func NewMediaType(mediaType string) (MediaType, error) {
	switch strings.ToLower(mediaType) {
	case "image", "gif", "video":
		return MediaType(mediaType), nil
	default:
		return "", ErrInvalidMediaType
	}
}
