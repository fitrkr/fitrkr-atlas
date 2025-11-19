package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries"
	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

func (h *Registry) CreateExercise(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.CreateExerciseCommand{}
	h.ExecuteCommand(cmd, true, w, r)
}

func (h *Registry) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.UpdateExerciseCommand{}
	h.ExecuteCommand(cmd, true, w, r)
}

func (h *Registry) DeleteExercise(w http.ResponseWriter, r *http.Request) {
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

func (h *Registry) GetExerciseByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := &queries.GetExerciseByIDQuery{
		ID: id,
	}

	h.ExecuteQuery(qry, w, r)
}

func (h *Registry) GetExerciseByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	qry := &queries.GetExerciseByNameQuery{
		Name: name,
	}

	h.ExecuteQuery(qry, w, r)
}
