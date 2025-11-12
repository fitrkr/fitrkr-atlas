package application

import (
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands/categories"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands/equipments"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands/exercises"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands/muscles"
)

type Commands struct {
	// Equipment
	CreateEquipment           *equipments.CreateEquipmentCommand
	UpdateEquipment           *equipments.UpdateEquipmentCommand
	DeleteEquipment           *equipments.DeleteEquipmentCommand
	CreateAttachment          *equipments.CreateAttachmentCommand
	UpdateAttachment          *equipments.UpdateAttachmentCommand
	DeleteAttachment          *equipments.DeleteAttachmentCommand
	CreateEquipmentAttachment *equipments.CreateEquipmentAttachmentCommand
	DeleteEquipmentAttachment *equipments.DeleteEquipmentAttachmentCommand

	// Muscle & MuscleGroup
	CreateMuscle *muscles.CreateMuscleCommand
	UpdateMuscle *muscles.UpdateMuscleCommand
	DeleteMuscle *muscles.DeleteMuscleCommand

	// Category & Subcategory
	CreateCategory *categories.CreateCategoryCommand
	UpdateCategory *categories.UpdateCategoryCommand
	DeleteCategory *categories.DeleteCategoryCommand

	// Exercise
	CreateExercise *exercises.CreateExerciseCommand
	UpdateExercise *exercises.UpdateExerciseCommand
	DeleteExercise *exercises.DeleteExerciseCommand
}
