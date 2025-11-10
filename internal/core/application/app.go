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
	muscleGroupWrite ports.MuscleGroupWrite,
	muscleGroupRead ports.MuscleGroupRead,
	muscleWrite ports.MuscleWrite,
	muscleRead ports.MuscleRead,

	// Category ports
	categoryWrite ports.CategoryWrite,
	categoryRead ports.CategoryRead,
	subcategoryWrite ports.SubcategoryWrite,
	subcategoryRead ports.SubcategoryRead,

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

	exerciseMediaWrite ports.ExerciseMediaWrite,
	exerciseMediaRead ports.ExerciseMediaRead,

	exerciseInstructionWrite ports.ExerciseInstructionWrite,
	exerciseInstructionRead ports.ExerciseInstructionRead,

	exerciseVariationWrite ports.ExerciseVariationWrite,
	// variationRead ports.ExerciseVariationRead,

	exerciseAlternateWrite ports.ExerciseAlternateWrite,
	// alternateRead ports.ExerciseAlternateRead,
) *Application {
	exerciseReadGroup := &ports.ExerciseReadGroup{
		ExerciseRead: exerciseRead,
		Alias:        exerciseAliasRead,
		Equipment:    exerciseEquipmentRead,
		Muscle:       exerciseMuscleRead,
		Category:     exerciseCategoryRead,
		Media:        exerciseMediaRead,
		Instruction:  exerciseInstructionRead,
	}

	exerciseWriteGroup := &ports.ExerciseWriteGroup{
		ExerciseWrite: exerciseWrite,
		Alias:         exerciseAliasWrite,
		Equipment:     exerciseEquipmentWrite,
		Muscle:        exerciseMuscleWrite,
		Category:      exerciseCategoryWrite,
		Media:         exerciseMediaWrite,
		Instruction:   exerciseInstructionWrite,
	}

	equipmentWriteGroup := &ports.EquipmentWriteGroup{
		EquipmentWrite: equipmentWrite,
		Attachment:     equipmentAttachmentWrite,
	}

	equipmentReadGroup := &ports.EquipmentReadGroup{
		EquipmentRead: equipmentRead,
		Attachment:    equipmentAttachmentRead,
	}

	muscleWriteGroup := &ports.MuscleWriteGroup{
		MuscleWrite: muscleWrite,
		Group:       muscleGroupWrite,
	}

	muscleReadGroup := &ports.MuscleReadGroup{
		MuscleRead: muscleRead,
		Group:      muscleGroupRead,
	}

	categoryWriteGroup := &ports.CategoryWriteGroup{
		CategoryWrite: categoryWrite,
		Subcategory:   subcategoryWrite,
	}

	categoryReadGroup := &ports.CategoryReadGroup{
		CategoryRead: categoryRead,
		Subcategory:  subcategoryRead,
	}

	write := &ports.Write{
		Exercise:   *exerciseWriteGroup,
		Equipment:  *equipmentWriteGroup,
		Attachment: attachmentWrite,
		Muscle:     *muscleWriteGroup,
		Category:   *categoryWriteGroup,
	}

	read := &ports.Read{
		Exercise:   *exerciseReadGroup,
		Equipment:  *equipmentReadGroup,
		Attachment: attachmentRead,
		Muscle:     *muscleReadGroup,
		Category:   *categoryReadGroup,
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
			CreateMuscleGroup: &muscles.CreateMuscleGroupCommand{
				Write: *write,
			},
			UpdateMuscleGroup: &muscles.UpdateMuscleGroupCommand{
				Write: *write,
				Read:  *read,
			},
			DeleteMuscleGroup: &muscles.DeleteMuscleGroupCommand{
				Write: *write,
			},
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
			CreateSubcategory: &categories.CreateSubcategoryCommand{
				Write: *write,
				Read:  *read,
			},
			UpdateSubcategory: &categories.UpdateSubcategoryCommand{
				Write: *write,
				Read:  *read,
			},
			DeleteSubcategory: &categories.DeleteSubcategoryCommand{
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
			GetMuscleGroupByID: &Qmuscles.GetMuscleGroupByIDQuery{
				Read: *read,
			},
			GetAllMuscleGroups: &Qmuscles.GetAllMuscleGroupsQuery{
				Read: *read,
			},
			GetMuscleByID: &Qmuscles.GetMuscleByIDQuery{
				Read: *read,
			},
			GetAllMuscles: &Qmuscles.GetAllMusclesQuery{
				Read: *read,
			},
			GetMusclesByMuscleGroupID: &Qmuscles.GetMusclesByMuscleGroupIDQuery{
				Read: *read,
			},

			// Category & Subcategory queries
			GetCategoryByID: &Qcategories.GetCategoryByIDQuery{
				Read: *read,
			},
			GetAllCategories: &Qcategories.GetAllCategoriesQuery{
				Read: *read,
			},
			GetSubcategoryByID: &Qcategories.GetSubcategoryByIDQuery{
				Read: *read,
			},
			GetAllSubcategories: &Qcategories.GetAllSubcategoriesQuery{
				Read: *read,
			},
			GetSubcategoriesByCategoryID: &Qcategories.GetSubcategoriesByCategoryIDQuery{
				Read: *read,
			},

			// Exercise queries
			GetExerciseAliasByID:       &Qexercises.GetAliasByIDQuery{Read: *read},
			GetExerciseCategoryByID:    &Qexercises.GetCategoryByIDQuery{Read: *read},
			GetExerciseEquipmentByID:   &Qexercises.GetEquipmentByIDQuery{Read: *read},
			GetExerciseMuscleByID:      &Qexercises.GetMuscleByIDQuery{Read: *read},
			GetExerciseByID:            &Qexercises.GetExerciseByIDQuery{Read: *read},
			GetExerciseByName:          &Qexercises.GetExerciseByNameQuery{Read: *read},
			GetExerciseInstructionByID: &Qexercises.GetInstructionByIDQuery{Read: *read},
			GetExerciseMediaByID:       &Qexercises.GetMediaByIDQuery{Read: *read},
		},
	}
}
