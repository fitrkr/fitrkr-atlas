// Package view
package view

import "time"

type View struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Type         string        `json:"type"`
	Difficulty   string        `json:"difficulty"`
	BodyPosition string        `json:"body_position"`
	Alias        []Alias       `json:"alias"`
	Muscle       []MuscleGroup `json:"muscle"`
	Equipment    []Equipment   `json:"equipment"`
	Category     []Category    `json:"category"`
	Media        []Media       `json:"media"`
	Instruction  []Instruction `json:"instruction"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    *time.Time    `json:"deleted_at"`
	PurgeAt      *time.Time    `json:"purge_at"`
}

func NewExerciseView(id int, name, description, exerciseType, difficulty, bodyposition string, muscle []MuscleGroup, equipment []Equipment, category []Category, createdAt, updatedAt time.Time, deletedAt, purgeAt *time.Time) View {
	return View{
		ID:           id,
		Name:         name,
		Description:  description,
		Type:         exerciseType,
		Difficulty:   difficulty,
		BodyPosition: bodyposition,
		Alias:        []Alias{},
		Muscle:       muscle,
		Equipment:    equipment,
		Category:     category,
		Media:        []Media{},
		Instruction:  []Instruction{},
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		DeletedAt:    deletedAt,
		PurgeAt:      purgeAt,
	}
}
