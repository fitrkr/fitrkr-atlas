// Package handlers
package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/application/mediator"
	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

type Registry struct {
	bus mediator.Bus
}

func NewHandlerRegistry(bus mediator.Bus) *Registry {
	return &Registry{
		bus: bus,
	}
}

type Handler interface {
	Handle(ctx context.Context) (any, error)
}

func (h *Registry) ExecuteCommand(cmd mediator.Handler, hasBody bool, w http.ResponseWriter, r *http.Request) error {
	if hasBody {
		if err := json.NewDecoder(r.Body).Decode(cmd); err != nil {
			web.ClientError(w, http.StatusBadRequest)
			return err
		}
	}

	resp, err := h.bus.Send(r.Context(), cmd)
	if err != nil {
		web.ServerError(w, err)
		return err
	}

	web.Response(w, http.StatusCreated, resp)
	return nil
}

func (h *Registry) ExecuteQuery(qry Handler, w http.ResponseWriter, r *http.Request) error {
	resp, err := h.bus.Send(r.Context(), qry)
	if err != nil {
		web.ServerError(w, err)
		return err
	}

	web.Response(w, http.StatusOK, resp)
	return nil
}

func (h *Registry) DecodeCommand(w http.ResponseWriter, r *http.Request, req any) (any, error) {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return nil, err
	}

	return req, nil
}

// first one is the command chain head

func (h *Registry) GetResp(w http.ResponseWriter, r *http.Request, cmd mediator.Handler) (any, error) {
	if err := json.NewDecoder(r.Body).Decode(cmd); err != nil {
		web.ClientError(w, http.StatusBadRequest)
		return nil, err
	}

	resp, err := h.bus.Send(r.Context(), cmd)
	if err != nil {
		web.ServerError(w, err)
		return nil, err
	}

	return resp, nil
}
