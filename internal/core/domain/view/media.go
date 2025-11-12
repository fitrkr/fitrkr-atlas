package view

import "time"

type Media struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	Type      string    `json:"type"`
	Order     int       `json:"order"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
