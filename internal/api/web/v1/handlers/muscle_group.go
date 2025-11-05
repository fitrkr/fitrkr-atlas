package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

func (h *Registry) CreateMuscleGroup(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.CreateMuscleGroup
	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) UpdateMuscleGroup(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.UpdateMuscleGroup
	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) DeleteMuscleGroup(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}
	cmd := h.Command.DeleteMuscleGroup
	cmd.ID = id
	ExecuteCommandNoBody(cmd, w, r)
}

func (h *Registry) GetMuscleGroupByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := h.Query.GetMuscleGroupByID
	qry.ID = id

	ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAllMuscleGroups(w http.ResponseWriter, r *http.Request) {
	qry := h.Query.GetAllMuscleGroups

	ExecuteQuery(qry, w, r)
}
