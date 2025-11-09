package exercise

import (
	"errors"
	"time"
)

var (
	ErrInvalidName = errors.New("empty name")
	ErrEmptyURL    = errors.New("empty url")
	ErrEmptyOrder  = errors.New("empty display order")
)

type Media struct {
	ID         *int
	ExerciseID int
	URL        string    `json:"url"` // TODO add a URL validator
	Type       MediaType `json:"type"`
	Order      int       `json:"order"`
	IsPrimary  bool      `json:"is_primary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewMedia(exerciseID, order int, url string, mediaType MediaType, isPrimary bool) (Media, error) {
	if exerciseID < 0 {
		return Media{}, ErrEmptyExericiseID
	}
	if order < 1 {
		return Media{}, ErrEmptyOrder
	}
	if url == "" {
		return Media{}, ErrEmptyURL
	}
	return Media{
		ExerciseID: exerciseID,
		URL:        url,
		Type:       mediaType,
		Order:      order,
		IsPrimary:  isPrimary,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
