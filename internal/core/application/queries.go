package application

import (
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/categories"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/equipments"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/exercises"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries/muscles"
)

type Queries struct {
	// Equipment
	GetEquipmentByID                    *equipments.GetEquipmentByIDQuery
	GetAllEquipments                    *equipments.GetAllEquipmentsQuery
	GetAttachmentByID                   *equipments.GetAttachmentByIDQuery
	GetAllAttachments                   *equipments.GetAllAttachmentsQuery
	GetEquipmentAttachmentByEquipmentID *equipments.GetEquipmentAttachmentByEquipmentIDQuery

	// Muscle & MuscleGroup
	GetMuscleGroupByID        *muscles.GetMuscleGroupByIDQuery
	GetAllMuscleGroups        *muscles.GetAllMuscleGroupsQuery
	GetMuscleByID             *muscles.GetMuscleByIDQuery
	GetAllMuscles             *muscles.GetAllMusclesQuery
	GetMusclesByMuscleGroupID *muscles.GetMusclesByMuscleGroupIDQuery

	// Category & Subcategory
	GetCategoryByID              *categories.GetCategoryByIDQuery
	GetAllCategories             *categories.GetAllCategoriesQuery
	GetSubcategoryByID           *categories.GetSubcategoryByIDQuery
	GetAllSubcategories          *categories.GetAllSubcategoriesQuery
	GetSubcategoriesByCategoryID *categories.GetSubcategoriesByCategoryIDQuery

	// Exercise
	GetExerciseAliasByID       *exercises.GetAliasByIDQuery
	GetExerciseCategoryByID    *exercises.GetCategoryByIDQuery
	GetExerciseEquipmentByID   *exercises.GetEquipmentByIDQuery
	GetExerciseMuscleByID      *exercises.GetMuscleByIDQuery
	GetExerciseByID            *exercises.GetExerciseByIDQuery
	GetExerciseByName          *exercises.GetExerciseByNameQuery
	GetExerciseInstructionByID *exercises.GetInstructionByIDQuery
	GetExerciseMediaByID       *exercises.GetMediaByIDQuery
}
