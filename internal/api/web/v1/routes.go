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
		"/view":       SetupViewRoutes(registry),
		"/exercise":   SetupExerciseRoutes(registry),
		"/equipment":  SetupEquipmentRoutes(registry),
		"/attachment": SetupAttachmentRoutes(registry),
		"/muscle":     SetupMuscleRoutes(registry),
		"/category":   SetupCategoryRoutes(registry),
	}

	// Mount the versioned routes
	for path, handler := range routes {
		r.Mount(path, handler)
	}

	return r
}

func SetupEquipmentRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetAllEquipments)
	r.Get("/{id}", h.GetEquipmentByID)

	r.Post("/", h.CreateEquipment)
	r.Put("/", h.UpdateEquipment)
	r.Delete("/{id}", h.DeleteEquipment)

	r.Route("/attachment", func(r chi.Router) {
	})

	return r
}

func SetupAttachmentRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetAllAttachments)
	r.Get("/{id}", h.GetAttachmentByID)

	r.Post("/", h.CreateAttachment)
	r.Put("/", h.UpdateAttachment)
	r.Delete("/{id}", h.DeleteAttachment)

	return r
}

func SetupMuscleRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetAllMuscles)
	r.Get("/{id}", h.GetMuscleByID)
	r.Get("/group/{type}", h.GetMusclesByGroupType)

	r.Post("/", h.CreateMuscle)
	r.Put("/", h.UpdateMuscle)
	r.Delete("/{id}", h.DeleteMuscle)

	return r
}

func SetupCategoryRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetAllCategories)
	r.Get("/{id}", h.GetCategoryByID)
	r.Get("/type/{type}", h.GetCategoriesByType)

	r.Post("/", h.CreateCategory)
	r.Put("/", h.UpdateCategory)
	r.Delete("/{id}", h.DeleteCategory)

	return r
}

func SetupViewRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetAllView)
	r.Get("/{id}", h.GetViewByID)

	return r
}

func SetupExerciseRoutes(h *handlers.Registry) http.Handler {
	r := chi.NewRouter()

	r.Get("/{id}", h.GetExerciseByID)
	r.Get("/name/{name}", h.GetExerciseByName)

	r.Post("/", h.CreateExercise)
	r.Put("/", h.UpdateExercise)
	r.Delete("/{id}", h.DeleteExercise)

	return r
}
