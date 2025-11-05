package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

func (h *Registry) CreateEquipment(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.CreateEquipment
	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) UpdateEquipment(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.UpdateEquipment
	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) DeleteEquipment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}
	cmd := h.Command.DeleteEquipment
	cmd.ID = id
	ExecuteCommandNoBody(cmd, w, r)
}

func (h *Registry) GetEquipmentByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := h.Query.GetEquipmentByID
	qry.ID = id

	ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAllEquipments(w http.ResponseWriter, r *http.Request) {
	qry := h.Query.GetAllEquipments

	ExecuteQuery(qry, w, r)
}
