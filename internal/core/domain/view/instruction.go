package view

import "time"

type Instruction struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
