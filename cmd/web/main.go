// Package main
package main

import (
	"github.com/cheezecakee/logr"
	"github.com/joho/godotenv"

	"github.com/fitrkr/atlas/internal/api/web"
	"github.com/fitrkr/atlas/internal/core/application/commands"
	"github.com/fitrkr/atlas/internal/core/application/mediator"
	"github.com/fitrkr/atlas/internal/core/application/queries"
	"github.com/fitrkr/atlas/internal/infrastructure/db/postgres"
)

func main() {
	logr.Init(&logr.PlainTextFormatter{}, logr.LevelInfo, nil)

	if err := godotenv.Load(); err != nil {
		logr.Get().Errorf("No .env file found: %v", err)
	}

	db := postgres.NewPostgresConn()
	defer db.Close()

	provider := postgres.NewProvider(db)
	read, write := provider.CreatePorts()

	commands.Init(write, read)
	queries.Init(read)

	registry := mediator.NewRegistry()
	queries.RegisterAll(registry)
	commands.RegisterAll(registry)

	// Initialize with subdirectory packages
	httpServer := web.NewApp(registry, web.WithPort(8080))
	httpServer.Run()
}
