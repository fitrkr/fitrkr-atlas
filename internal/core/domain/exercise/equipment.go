package exercise

import (
	"errors"
	"time"
)

var ErrEmptyEquipmentID = errors.New("empty equipment id")

type ExerciseEquipment struct {
	ID           int       `json:"id"`
	EquipmentID  int       `json:"equipment_id"`
	AttachmentID *int      `json:"attachment_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewExerciseEquipment(exerciseID, equipmentID int, attachmentID *int) (ExerciseEquipment, error) {
	if exerciseID < 0 {
		return ExerciseEquipment{}, ErrEmptyExericiseID
	}
	if equipmentID < 0 {
		return ExerciseEquipment{}, ErrEmptyEquipmentID
	}
	return ExerciseEquipment{
		ID:           exerciseID,
		EquipmentID:  equipmentID,
		AttachmentID: attachmentID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}
