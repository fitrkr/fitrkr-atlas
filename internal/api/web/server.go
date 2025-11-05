// Package web
package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	v1 "github.com/cheezecakee/fitrkr-atlas/internal/api/web/v1"
	"github.com/cheezecakee/fitrkr-atlas/internal/api/web/v1/handlers"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application"
)

type App struct {
	chi      *chi.Mux
	port     int
	registry *handlers.Registry
}

func NewApp(application *application.Application, opts ...AppOption) *App {
	app := &App{
		port: 8000,
		chi:  chi.NewRouter(),
	}

	for _, applyOption := range opts {
		applyOption(app)
	}

	app.registry = handlers.NewHandlerRegistry(application)

	fs := http.FileServer(http.Dir("internal/api/web/docs"))

	app.chi.Handle("/api/v1/docs/*", http.StripPrefix("/api/v1/docs/", fs))

	app.chi.Mount("/api/v1", v1.RegisterRoutes(app.registry))

	return app
}

func (a *App) Run() error {
	addr := fmt.Sprintf(":%d", a.port)
	return http.ListenAndServe(addr, a.chi)
}
