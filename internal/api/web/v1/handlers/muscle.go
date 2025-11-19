package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries"
	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

func (h *Registry) CreateMuscle(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.CreateMuscleCommand{}
	h.ExecuteCommand(cmd, true, w, r)
}

func (h *Registry) UpdateMuscle(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.UpdateMuscleCommand{}
	h.ExecuteCommand(cmd, true, w, r)
}

func (h *Registry) DeleteMuscle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}
	cmd := &commands.DeleteMuscleCommand{
		ID: id,
	}
	h.ExecuteCommand(cmd, false, w, r)
}

func (h *Registry) GetMuscleByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := &queries.GetMuscleByIDQuery{
		ID: id,
	}

	h.ExecuteQuery(qry, w, r)
}

func (h *Registry) GetMusclesByGroupType(w http.ResponseWriter, r *http.Request) {
	groupType := chi.URLParam(r, "type")

	qry := &queries.GetMusclesByGroupTypeQuery{
		GroupType: groupType,
	}

	h.ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAllMuscles(w http.ResponseWriter, r *http.Request) {
	qry := &queries.GetAllMusclesQuery{}
	h.ExecuteQuery(qry, w, r)
}
