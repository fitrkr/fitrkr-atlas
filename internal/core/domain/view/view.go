// Package view
package view

import "time"

type View struct {
	ID          int           `json:"id"` // This is the same as exercise id
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Difficulty  string        `json:"difficulty"`
	Position    string        `json:"position"`
	Alias       []string      `json:"alias"`
	Equipment   Equipment     `json:"equipment"`
	MuscleGroup []MuscleGroup `json:"muscle_group"`
	Category    []Category    `json:"category"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   *time.Time    `json:"deleted_at"`
	PurgeAt     *time.Time    `json:"purge_at"`
}

func NewExerciseView(exerciseID int, name, description, difficulty, position string, equipment Equipment, muscleGroup []MuscleGroup, category []Category, createdAt, updatedAt time.Time, deletedAt, purgeAt *time.Time, alias *[]string) View {
	return View{
		ID:          exerciseID,
		Name:        name,
		Description: description,
		Difficulty:  difficulty,
		Position:    position,
		Alias:       *alias,
		MuscleGroup: muscleGroup,
		Equipment:   equipment,
		Category:    category,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
		PurgeAt:     purgeAt,
	}
}
