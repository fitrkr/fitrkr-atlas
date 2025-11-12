// Package application
package application

import (
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands/categories"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands/equipments"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands/exercises"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands/muscles"
	Qcategories "github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/categories"
	Qequipments "github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/equipments"
	Qexercises "github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/exercises"
	Qmuscles "github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/muscles"
	Qview "github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/view"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func New(
	// Equipment ports
	equipmentWrite ports.EquipmentWrite,
	equipmentRead ports.EquipmentRead,
	equipmentAttachmentWrite ports.EquipmentAttachmentWrite,
	equipmentAttachmentRead ports.EquipmentAttachmentRead,
	attachmentWrite ports.AttachmentWrite,
	attachmentRead ports.AttachmentRead,

	// Muscle ports
	muscleWrite ports.MuscleWrite,
	muscleRead ports.MuscleRead,

	// Category ports
	categoryWrite ports.CategoryWrite,
	categoryRead ports.CategoryRead,

	// Exercise ports
	exerciseWrite ports.ExerciseWrite,
	exerciseRead ports.ExerciseRead,

	exerciseAliasWrite ports.ExerciseAliasWrite,
	exerciseAliasRead ports.ExerciseAliasRead,

	exerciseEquipmentWrite ports.ExerciseEquipmentWrite,
	exerciseEquipmentRead ports.ExerciseEquipmentRead,

	exerciseMuscleWrite ports.ExerciseMuscleWrite,
	exerciseMuscleRead ports.ExerciseMuscleRead,

	exerciseCategoryWrite ports.ExerciseCategoryWrite,
	exerciseCategoryRead ports.ExerciseCategoryRead,

	viewWrite ports.ViewWrite,
	viewRead ports.ViewRead,
) *Application {
	exerciseReadGroup := &ports.ExerciseReadGroup{
		ExerciseRead: exerciseRead,
		Alias:        exerciseAliasRead,
		Equipment:    exerciseEquipmentRead,
		Muscle:       exerciseMuscleRead,
		Category:     exerciseCategoryRead,
	}

	exerciseWriteGroup := &ports.ExerciseWriteGroup{
		ExerciseWrite: exerciseWrite,
		Alias:         exerciseAliasWrite,
		Equipment:     exerciseEquipmentWrite,
		Muscle:        exerciseMuscleWrite,
		Category:      exerciseCategoryWrite,
	}

	equipmentWriteGroup := &ports.EquipmentWriteGroup{
		EquipmentWrite: equipmentWrite,
		Attachment:     equipmentAttachmentWrite,
	}

	equipmentReadGroup := &ports.EquipmentReadGroup{
		EquipmentRead: equipmentRead,
		Attachment:    equipmentAttachmentRead,
	}

	write := &ports.Write{
		Exercise:   *exerciseWriteGroup,
		Equipment:  *equipmentWriteGroup,
		Attachment: attachmentWrite,
		Muscle:     muscleWrite,
		Category:   categoryWrite,
		View:       viewWrite,
	}

	read := &ports.Read{
		Exercise:   *exerciseReadGroup,
		Equipment:  *equipmentReadGroup,
		Attachment: attachmentRead,
		Muscle:     muscleRead,
		Category:   categoryRead,
		View:       viewRead,
	}

	return &Application{
		Commands: Commands{
			// Exercise commands
			CreateExercise: &exercises.CreateExerciseCommand{
				Write: *write,
				Read:  *read,
			},
			UpdateExercise: &exercises.UpdateExerciseCommand{
				Write: *write,
				Read:  *read,
			},
			DeleteExercise: &exercises.DeleteExerciseCommand{
				Write: *write,
				Read:  *read,
			},

			// Equipment commands
			CreateEquipment: &equipments.CreateEquipmentCommand{
				Write: *write,
				Read:  *read,
			},
			UpdateEquipment: &equipments.UpdateEquipmentCommand{
				Write: *write,
				Read:  *read,
			},
			DeleteEquipment: &equipments.DeleteEquipmentCommand{
				Write: *write,
			},
			CreateAttachment: &equipments.CreateAttachmentCommand{
				Write: *write,
				Read:  *read,
			},
			UpdateAttachment: &equipments.UpdateAttachmentCommand{
				Write: *write,
				Read:  *read,
			},
			DeleteAttachment: &equipments.DeleteAttachmentCommand{
				Write: *write,
			},
			CreateEquipmentAttachment: &equipments.CreateEquipmentAttachmentCommand{
				Write: *write,
				Read:  *read,
			},
			DeleteEquipmentAttachment: &equipments.DeleteEquipmentAttachmentCommand{
				Write: *write,
			},
			// Muscle & MuscleGroup commands
			CreateMuscle: &muscles.CreateMuscleCommand{
				Write: *write,
				Read:  *read,
			},
			UpdateMuscle: &muscles.UpdateMuscleCommand{
				Write: *write,
				Read:  *read,
			},
			DeleteMuscle: &muscles.DeleteMuscleCommand{
				Write: *write,
			},
			// Category & Subcategory commands
			CreateCategory: &categories.CreateCategoryCommand{
				Write: *write,
			},
			UpdateCategory: &categories.UpdateCategoryCommand{
				Write: *write,
				Read:  *read,
			},
			DeleteCategory: &categories.DeleteCategoryCommand{
				Write: *write,
			},
		},
		Queries: Queries{
			// Equipment queries
			GetEquipmentByID: &Qequipments.GetEquipmentByIDQuery{
				Read: *read,
			},
			GetAllEquipments: &Qequipments.GetAllEquipmentsQuery{
				Read: *read,
			},
			GetAttachmentByID: &Qequipments.GetAttachmentByIDQuery{
				Read: *read,
			},
			GetAllAttachments: &Qequipments.GetAllAttachmentsQuery{
				Read: *read,
			},
			GetEquipmentAttachmentByEquipmentID: &Qequipments.GetEquipmentAttachmentByEquipmentIDQuery{
				Read: *read,
			},

			// Muscle & MuscleGroup queries
			GetMuscleByID: &Qmuscles.GetMuscleByIDQuery{
				Read: *read,
			},
			GetAllMuscles: &Qmuscles.GetAllMusclesQuery{
				Read: *read,
			},

			// Category & Subcategory queries
			GetCategoryByID: &Qcategories.GetCategoryByIDQuery{
				Read: *read,
			},
			GetAllCategories: &Qcategories.GetAllCategoriesQuery{
				Read: *read,
			},

			// Exercise queries
			GetExerciseAliasByID:     &Qexercises.GetAliasByIDQuery{Read: *read},
			GetExerciseCategoryByID:  &Qexercises.GetCategoryByIDQuery{Read: *read},
			GetExerciseEquipmentByID: &Qexercises.GetEquipmentByIDQuery{Read: *read},
			GetExerciseMuscleByID:    &Qexercises.GetMuscleByIDQuery{Read: *read},
			GetExerciseByID:          &Qexercises.GetExerciseByIDQuery{Read: *read},
			GetExerciseByName:        &Qexercises.GetExerciseByNameQuery{Read: *read},

			// View queries
			GetViewByID: &Qview.GetViewByIDQuery{Read: *read},
			GetAllView:  &Qview.GetAllViewQuery{Read: *read},
		},
	}
}
