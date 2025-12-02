package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/application/queries"
	"github.com/cheezecakee/fitrkr/atlas/pkg/web"
)

func (h *Registry) GetViewByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return
	}

	qry := &queries.GetViewByIDQuery{
		ID: id,
	}

	h.ExecuteQuery(qry, w, r)
}

func (h *Registry) GetAllView(w http.ResponseWriter, r *http.Request) {
	qry := &queries.GetAllViewQuery{}

	h.ExecuteQuery(qry, w, r)
}
