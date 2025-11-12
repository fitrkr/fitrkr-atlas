package application

import (
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/categories"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/equipments"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/exercises"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/muscles"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/view"
)

type Queries struct {
	// Equipment
	GetEquipmentByID                    *equipments.GetEquipmentByIDQuery
	GetAllEquipments                    *equipments.GetAllEquipmentsQuery
	GetAttachmentByID                   *equipments.GetAttachmentByIDQuery
	GetAllAttachments                   *equipments.GetAllAttachmentsQuery
	GetEquipmentAttachmentByEquipmentID *equipments.GetEquipmentAttachmentByEquipmentIDQuery

	// Muscle & MuscleGroup
	GetMuscleByID         *muscles.GetMuscleByIDQuery
	GetAllMuscles         *muscles.GetAllMusclesQuery
	GetMusclesByGroupType *muscles.GetMusclesByMuscleGroupIDQuery

	// Category & Subcategory
	GetCategoryByID  *categories.GetCategoryByIDQuery
	GetAllCategories *categories.GetAllCategoriesQuery

	// Exercise
	GetExerciseAliasByID     *exercises.GetAliasByIDQuery
	GetExerciseCategoryByID  *exercises.GetCategoryByIDQuery
	GetExerciseEquipmentByID *exercises.GetEquipmentByIDQuery
	GetExerciseMuscleByID    *exercises.GetMuscleByIDQuery
	GetExerciseByID          *exercises.GetExerciseByIDQuery
	GetExerciseByName        *exercises.GetExerciseByNameQuery

	// View
	GetViewByID *view.GetViewByIDQuery
	GetAllView  *view.GetAllViewQuery
}
