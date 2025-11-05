// Package handlers
package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/application"
	"github.com/cheezecakee/fitrkr-atlas/pkg/web"
)

type Registry struct {
	Command application.Commands
	Query   application.Queries
}

func NewHandlerRegistry(app *application.Application) *Registry {
	return &Registry{
		Command: app.Commands,
		Query:   app.Queries,
	}
}

type Handler interface {
	Handle(ctx context.Context) (any, error)
}

func ExecuteCommand(cmd Handler, hasBody bool, w http.ResponseWriter, r *http.Request) error {
	if hasBody {
		if err := json.NewDecoder(r.Body).Decode(cmd); err != nil {
			web.ClientError(w, http.StatusBadRequest)
			logr.Get().Debugf("failed to execute command: %s", err)
			return err
		}
	}

	resp, err := cmd.Handle(r.Context())
	if err != nil {
		web.ServerError(w, err)
		return err
	}

	web.Response(w, http.StatusCreated, resp)
	return nil
}

func ExecuteCommandWithBody(cmd Handler, w http.ResponseWriter, r *http.Request) error {
	return ExecuteCommand(cmd, true, w, r)
}

func ExecuteCommandNoBody(cmd Handler, w http.ResponseWriter, r *http.Request) error {
	return ExecuteCommand(cmd, false, w, r)
}

func ExecuteQuery(qry Handler, w http.ResponseWriter, r *http.Request) error {
	resp, err := qry.Handle(r.Context())
	if err != nil {
		web.ServerError(w, err)
		return err
	}

	web.Response(w, http.StatusOK, resp)
	return nil
}
