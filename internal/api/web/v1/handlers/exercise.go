package handlers

import (
	"net/http"
	"strconv"

	"github.com/cheezecakee/logr"
	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries"
	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

func (h *Registry) CreateExercise(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.CreateExerciseCommand{}
	resp, err := h.GetResp(w, r, cmd)
	if err != nil {
		web.ServerError(w, err)
		return
	}

	exerciseID := resp.(commands.CreateExerciseResp).ID

	cmdView := &commands.BuildViewCommand{ExerciseID: exerciseID, Create: true}

	h.ExecuteCommand(cmdView, false, w, r)
}

func (h *Registry) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.UpdateExerciseCommand{}
	resp, err := h.GetResp(w, r, cmd)
	if err != nil {
		web.ServerError(w, err)
		return
	}

	exercise := resp.(commands.UpdateExerciseResp).Exercise
	logr.Get().Infof("updated Exercise: %v", exercise)

	cmdView := &commands.BuildViewCommand{ExerciseID: *exercise.ID, Create: false}

	h.ExecuteCommand(cmdView, false, w, r)
}

func (h *Registry) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	cmd := &commands.DeleteExerciseCommand{
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
