// Package handlers
package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cheezecakee/logr"

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
			logr.Get().Debugf("failed to decode command: %s", err)
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

func (h *Registry) CommandChain(w http.ResponseWriter, r *http.Request, cmds []mediator.Handler) error {
	for _, cmd := range cmds {
		_, err := h.bus.Send(r.Context(), cmd)
		if err != nil {
			web.ServerError(w, err)
			return err
		}
	}
	return nil
}

func (h *Registry) ExecuteCommandSequence(ctx context.Context, cmds []mediator.Handler) (lastResp any, err error) {
	for _, cmd := range cmds {
		lastResp, err = h.bus.Send(ctx, cmd)
		if err != nil {
			return nil, err
		}
	}
	return lastResp, nil
}
