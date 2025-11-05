// Package application
package application

import (
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands/categories"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands/equipments"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands/muscles"
	Qcategories "github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/categories"
	Qequipments "github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/equipments"
	Qmuscles "github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/muscles"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	// Equipment
	CreateEquipment  *equipments.CreateEquipmentCommand
	UpdateEquipment  *equipments.UpdateEquipmentCommand
	DeleteEquipment  *equipments.DeleteEquipmentCommand
	CreateAttachment *equipments.CreateAttachmentCommand
	UpdateAttachment *equipments.UpdateAttachmentCommand
	DeleteAttachment *equipments.DeleteAttachmentCommand

	// Muscle & MuscleGroup
	CreateMuscleGroup *muscles.CreateMuscleGroupCommand
	UpdateMuscleGroup *muscles.UpdateMuscleGroupCommand
	DeleteMuscleGroup *muscles.DeleteMuscleGroupCommand
	CreateMuscle      *muscles.CreateMuscleCommand
	UpdateMuscle      *muscles.UpdateMuscleCommand
	DeleteMuscle      *muscles.DeleteMuscleCommand

	// Category & Subcategory
	CreateCategory    *categories.CreateCategoryCommand
	UpdateCategory    *categories.UpdateCategoryCommand
	DeleteCategory    *categories.DeleteCategoryCommand
	CreateSubcategory *categories.CreateSubcategoryCommand
	UpdateSubcategory *categories.UpdateSubcategoryCommand
	DeleteSubcategory *categories.DeleteSubcategoryCommand
}

type Queries struct {
	// Equipment
	GetEquipmentByID           *Qequipments.GetEquipmentByIDQuery
	GetAllEquipments           *Qequipments.GetAllEquipmentsQuery
	GetAttachmentByID          *Qequipments.GetAttachmentByIDQuery
	GetAllAttachments          *Qequipments.GetAllAttachmentsQuery
	GetAttachmentByEquipmentID *Qequipments.GetAttachmentsByEquipmentIDQuery

	// Muscle & MuscleGroup
	GetMuscleGroupByID        *Qmuscles.GetMuscleGroupByIDQuery
	GetAllMuscleGroups        *Qmuscles.GetAllMuscleGroupsQuery
	GetMuscleByID             *Qmuscles.GetMuscleByIDQuery
	GetAllMuscles             *Qmuscles.GetAllMusclesQuery
	GetMusclesByMuscleGroupID *Qmuscles.GetMusclesByMuscleGroupIDQuery

	// Category & Subcategory
	GetCategoryByID              *Qcategories.GetCategoryByIDQuery
	GetAllCategories             *Qcategories.GetAllCategoriesQuery
	GetSubcategoryByID           *Qcategories.GetSubcategoryByIDQuery
	GetAllSubcategories          *Qcategories.GetAllSubcategoriesQuery
	GetSubcategoriesByCategoryID *Qcategories.GetSubcategoriesByCategoryIDQuery
}

func New(
	// Equipment ports
	equipmentWrite ports.EquipmentWrite,
	equipmentRead ports.EquipmentRead,
	attachmentWrite ports.EquipmentAttachmentWrite,
	attachmentRead ports.EquipmentAttachmentRead,

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
) *Application {
	return &Application{
		Commands: Commands{
			// Equipment commands
			CreateEquipment: &equipments.CreateEquipmentCommand{
				Write: equipmentWrite,
				Read:  equipmentRead,
			},
			UpdateEquipment: &equipments.UpdateEquipmentCommand{
				Write: equipmentWrite,
				Read:  equipmentRead,
			},
			DeleteEquipment: &equipments.DeleteEquipmentCommand{
				Write: equipmentWrite,
			},
			CreateAttachment: &equipments.CreateAttachmentCommand{
				Write: attachmentWrite,
				Read:  equipmentRead,
			},
			UpdateAttachment: &equipments.UpdateAttachmentCommand{
				Write:         attachmentWrite,
				Read:          attachmentRead,
				ReadEquipment: equipmentRead,
			},
			DeleteAttachment: &equipments.DeleteAttachmentCommand{
				Write: attachmentWrite,
			},

			// Muscle & MuscleGroup commands
			CreateMuscleGroup: &muscles.CreateMuscleGroupCommand{
				Write: muscleGroupWrite,
			},
			UpdateMuscleGroup: &muscles.UpdateMuscleGroupCommand{
				Write: muscleGroupWrite,
				Read:  muscleGroupRead,
			},
			DeleteMuscleGroup: &muscles.DeleteMuscleGroupCommand{
				Write: muscleGroupWrite,
			},
			CreateMuscle: &muscles.CreateMuscleCommand{
				Write: muscleWrite,
				Read:  muscleGroupRead,
			},
			UpdateMuscle: &muscles.UpdateMuscleCommand{
				Write:     muscleWrite,
				Read:      muscleRead,
				ReadGroup: muscleGroupRead,
			},
			DeleteMuscle: &muscles.DeleteMuscleCommand{
				Write: muscleWrite,
			},

			// Category & Subcategory commands
			CreateCategory: &categories.CreateCategoryCommand{
				Write: categoryWrite,
			},
			UpdateCategory: &categories.UpdateCategoryCommand{
				Write: categoryWrite,
				Read:  categoryRead,
			},
			DeleteCategory: &categories.DeleteCategoryCommand{
				Write: categoryWrite,
			},
			CreateSubcategory: &categories.CreateSubcategoryCommand{
				Write: subcategoryWrite,
				Read:  categoryRead,
			},
			UpdateSubcategory: &categories.UpdateSubcategoryCommand{
				Write:        subcategoryWrite,
				Read:         subcategoryRead,
				ReadCategory: categoryRead,
			},
			DeleteSubcategory: &categories.DeleteSubcategoryCommand{
				Write: subcategoryWrite,
			},
		},
		Queries: Queries{
			// Equipment queries
			GetEquipmentByID: &Qequipments.GetEquipmentByIDQuery{
				Read: equipmentRead,
			},
			GetAllEquipments: &Qequipments.GetAllEquipmentsQuery{
				Read: equipmentRead,
			},
			GetAttachmentByID: &Qequipments.GetAttachmentByIDQuery{
				Read: attachmentRead,
			},
			GetAllAttachments: &Qequipments.GetAllAttachmentsQuery{
				Read: attachmentRead,
			},
			GetAttachmentByEquipmentID: &Qequipments.GetAttachmentsByEquipmentIDQuery{
				Read: attachmentRead,
			},

			// Muscle & MuscleGroup queries
			GetMuscleGroupByID: &Qmuscles.GetMuscleGroupByIDQuery{
				Read: muscleGroupRead,
			},
			GetAllMuscleGroups: &Qmuscles.GetAllMuscleGroupsQuery{
				Read: muscleGroupRead,
			},
			GetMuscleByID: &Qmuscles.GetMuscleByIDQuery{
				Read: muscleRead,
			},
			GetAllMuscles: &Qmuscles.GetAllMusclesQuery{
				Read: muscleRead,
			},
			GetMusclesByMuscleGroupID: &Qmuscles.GetMusclesByMuscleGroupIDQuery{
				Read: muscleRead,
			},

			// Category & Subcategory queries
			GetCategoryByID: &Qcategories.GetCategoryByIDQuery{
				Read: categoryRead,
			},
			GetAllCategories: &Qcategories.GetAllCategoriesQuery{
				Read: categoryRead,
			},
			GetSubcategoryByID: &Qcategories.GetSubcategoryByIDQuery{
				Read: subcategoryRead,
			},
			GetAllSubcategories: &Qcategories.GetAllSubcategoriesQuery{
				Read: subcategoryRead,
			},
			GetSubcategoriesByCategoryID: &Qcategories.GetSubcategoriesByCategoryIDQuery{
				Read: subcategoryRead,
			},
		},
	}
}
