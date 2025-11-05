package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

func (h *Registry) CreateAttachment(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.CreateAttachment

	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) UpdateAttachment(w http.ResponseWriter, r *http.Request) {
	cmd := h.Command.UpdateAttachment

	ExecuteCommandWithBody(cmd, w, r)
}

func (h *Registry) DeleteAttachment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}
	cmd := h.Command.DeleteAttachment
	cmd.ID = id

	ExecuteCommandNoBody(cmd, w, r)
}

func (h *Registry) GetAttachmentByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := h.Query.GetAttachmentByID
	qry.ID = id

	ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAttachmentsByEquipmentID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "equipmentID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := h.Query.GetAttachmentByEquipmentID
	qry.EquipmentID = id

	ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAllAttachments(w http.ResponseWriter, r *http.Request) {
	qry := h.Query.GetAllAttachments

	ExecuteQuery(qry, w, r)
}
