// Package v1
package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/internal/api/web/v1/handlers"
)

func RegisterRoutes(registry *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	routes := map[string]http.Handler{
		"/equipment":    SetupEquipmentRoutes(registry),
		"/attachment":   SetupAttachmentRoutes(registry),
		"/muscle-group": SetupMuscleGroupRoutes(registry),
		"/muscle":       SetupMuscleRoutes(registry),
		"/category":     SetupCategoryRoutes(registry),
		"/subcategory":  SetupSubcategoryRoutes(registry),
	}

	// Mount the versioned routes
	for path, handler := range routes {
		r.Mount(path, handler)
	}

	return r
}

func SetupEquipmentRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/all", h.GetAllEquipments)
	r.Get("/{id}", h.GetEquipmentByID)

	r.Post("/", h.CreateEquipment)
	r.Put("/", h.UpdateEquipment)
	r.Delete("/{id}", h.DeleteEquipment)

	return r
}

func SetupAttachmentRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/all", h.GetAllAttachments)
	r.Get("/{id}", h.GetAttachmentByID)
	r.Get("/equipment/{equipmentID}", h.GetAttachmentsByEquipmentID)

	r.Post("/", h.CreateAttachment)
	r.Put("/", h.UpdateAttachment)
	r.Delete("/{id}", h.DeleteAttachment)

	return r
}

func SetupMuscleGroupRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/all", h.GetAllMuscleGroups)
	r.Get("/{id}", h.GetMuscleGroupByID)

	r.Post("/", h.CreateMuscleGroup)
	r.Put("/", h.UpdateMuscleGroup)
	r.Delete("/{id}", h.DeleteMuscleGroup)

	return r
}

func SetupMuscleRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/all", h.GetAllMuscles)
	r.Get("/{id}", h.GetMuscleByID)
	r.Get("/group/{muscleGroupID}", h.GetMusclesByMuscleGroupID)

	r.Post("/", h.CreateMuscle)
	r.Put("/", h.UpdateMuscle)
	r.Delete("/{id}", h.DeleteMuscle)

	return r
}

func SetupCategoryRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/all", h.GetAllCategories)
	r.Get("/{id}", h.GetCategoryByID)

	r.Post("/", h.CreateCategory)
	r.Put("/", h.UpdateCategory)
	r.Delete("/{id}", h.DeleteCategory)

	return r
}

func SetupSubcategoryRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/all", h.GetAllSubcategories)
	r.Get("/{id}", h.GetSubcategoryByID)
	r.Get("/category/{categoryID}", h.GetSubcategoriesByCategoryID)

	r.Post("/", h.CreateSubcategory)
	r.Put("/", h.UpdateSubcategory)
	r.Delete("/{id}", h.DeleteSubcategory)

	return r
}
