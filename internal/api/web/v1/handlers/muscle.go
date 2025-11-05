package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

func (h *Registry) CreateMuscle(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.CreateMuscle

	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) UpdateMuscle(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.UpdateMuscle

	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) DeleteMuscle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}
	cmd := h.Command.DeleteMuscle
	cmd.ID = id

	ExecuteCommandNoBody(cmd, w, r)
}

func (h *Registry) GetMuscleByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := h.Query.GetMuscleByID
	qry.ID = id

	ExecuteQuery(qry, w, r)
}

func (h *Registry) GetMusclesByMuscleGroupID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "muscleGroupID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := h.Query.GetMusclesByMuscleGroupID
	qry.MuscleGroupID = id

	ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAllMuscles(w http.ResponseWriter, r *http.Request) {
	qry := h.Query.GetAllMuscles

	ExecuteQuery(qry, w, r)
}
