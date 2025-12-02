package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/application/commands"
	"github.com/cheezecakee/fitrkr/atlas/internal/core/application/queries"
	"github.com/cheezecakee/fitrkr/atlas/pkg/web"
)

func (h *Registry) CreateCategory(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.CreateCategoryCommand{}
	h.ExecuteCommand(cmd, true, w, r)
}

func (h *Registry) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.UpdateCategoryCommand{}
	h.ExecuteCommand(cmd, true, w, r)
}

func (h *Registry) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}
	cmd := &commands.DeleteCategoryCommand{
		ID: id,
	}
	h.ExecuteCommand(cmd, false, w, r)
}

func (h *Registry) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := &queries.GetCategoryByIDQuery{
		ID: id,
	}

	h.ExecuteQuery(qry, w, r)
}

func (h *Registry) GetCategoriesByType(w http.ResponseWriter, r *http.Request) {
	groupType := chi.URLParam(r, "type")

	qry := &queries.GetCategoriesByTypeQuery{
		Type: groupType,
	}

	h.ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	qry := &queries.GetAllCategoriesQuery{}

	h.ExecuteQuery(qry, w, r)
}
