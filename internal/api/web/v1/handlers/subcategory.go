package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

func (h *Registry) CreateSubcategory(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.CreateSubcategory

	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) UpdateSubcategory(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.UpdateSubcategory

	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) DeleteSubcategory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}
	cmd := h.Command.DeleteSubcategory
	cmd.ID = id
	ExecuteCommandNoBody(cmd, w, r)
}

func (h *Registry) GetSubcategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := h.Query.GetSubcategoryByID
	qry.ID = id

	ExecuteQuery(qry, w, r)
}

func (h *Registry) GetSubcategoriesByCategoryID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "categoryID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := h.Query.GetSubcategoriesByCategoryID
	qry.CategoryID = id

	ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAllSubcategories(w http.ResponseWriter, r *http.Request) {
	qry := h.Query.GetAllSubcategories

	ExecuteQuery(qry, w, r)
}
