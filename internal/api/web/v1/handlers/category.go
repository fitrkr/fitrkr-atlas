package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

func (h *Registry) CreateCategory(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.CreateCategory
	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.UpdateCategory
	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}
	cmd := h.Command.DeleteCategory
	cmd.ID = id
	ExecuteCommandNoBody(cmd, w, r)
}

func (h *Registry) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := h.Query.GetCategoryByID
	qry.ID = id

	ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	qry := h.Query.GetAllCategories

	ExecuteQuery(qry, w, r)
}
