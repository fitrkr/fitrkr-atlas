package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries"
	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

func (h *Registry) CreateEquipment(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.CreateEquipmentCommand{}
	h.ExecuteCommand(cmd, true, w, r)
}

func (h *Registry) CreateEquipmentAttachment(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.CreateEquipmentAttachmentCommand{}
	h.ExecuteCommand(cmd, true, w, r)
}

func (h *Registry) UpdateEquipment(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.UpdateEquipmentCommand{}
	h.ExecuteCommand(cmd, true, w, r)
}

func (h *Registry) DeleteEquipment(w http.ResponseWriter, r *http.Request) {
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

func (h *Registry) GetEquipmentByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := &queries.GetEquipmentByIDQuery{
		ID: id,
	}

	h.ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAllEquipments(w http.ResponseWriter, r *http.Request) {
	qry := &queries.GetAllEquipmentsQuery{}

	h.ExecuteQuery(qry, w, r)
}
