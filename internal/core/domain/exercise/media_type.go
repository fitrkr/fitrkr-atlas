package exercise

import (
	"errors"
	"strings"
)

type MediaType int

const (
	MediaImage MediaType = iota
	MediaGif
	MediaVideo
)

var ErrInvalidMediaType = errors.New("invalid media type")

func NewMediaType(mediaType string) (MediaType, error) {
	switch strings.ToLower(mediaType) {
	case "image":
		return MediaImage, nil
	case "gif":
		return MediaGif, nil
	case "video":
		return MediaVideo, nil
	default:
		return MediaImage, ErrInvalidMediaType
	}
}

func (m MediaType) ToString() string {
	switch m {
	case MediaImage:
		return "image"
	case MediaGif:
		return "gif"
	case MediaVideo:
		return "video"
	default:
		return "unknown"
	}
}
