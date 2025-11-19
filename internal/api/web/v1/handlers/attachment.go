package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/commands"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/queries"
	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

func (h *Registry) CreateAttachment(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.CreateAttachmentCommand{}
	h.ExecuteCommand(cmd, true, w, r)
}

func (h *Registry) UpdateAttachment(w http.ResponseWriter, r *http.Request) {
	cmd := &commands.UpdateAttachmentCommand{}
	h.ExecuteCommand(cmd, true, w, r)
}

func (h *Registry) DeleteAttachment(w http.ResponseWriter, r *http.Request) {
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

func (h *Registry) GetAttachmentByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := &queries.GetAttachmentByIDQuery{
		ID: id,
	}

	h.ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAllAttachments(w http.ResponseWriter, r *http.Request) {
	qry := &queries.GetAllAttachmentsQuery{}

	h.ExecuteQuery(qry, w, r)
}
