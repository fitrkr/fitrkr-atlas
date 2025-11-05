// Package main
package main

import (
	"github.com/cheezecakee/logr"
	"github.com/joho/godotenv"

	"github.com/cheezecakee/fitrkr-atlas/internal/api/web"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/application"
	"github.com/cheezecakee/fitrkr-atlas/internal/infrastructure/db/postgres"
	"github.com/cheezecakee/fitrkr-atlas/internal/infrastructure/db/postgres/attachment"
	"github.com/cheezecakee/fitrkr-atlas/internal/infrastructure/db/postgres/category"
	"github.com/cheezecakee/fitrkr-atlas/internal/infrastructure/db/postgres/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/infrastructure/db/postgres/muscle"
	"github.com/cheezecakee/fitrkr-atlas/internal/infrastructure/db/postgres/muscle_group"
	"github.com/cheezecakee/fitrkr-atlas/internal/infrastructure/db/postgres/subcategory"
)

func main() {
	logr.Init(&logr.PlainTextFormatter{}, logr.LevelInfo, nil)

	if err := godotenv.Load(); err != nil {
		logr.Get().Errorf("No .env file found: %v", err)
	}

	db := postgres.NewPostgresConn()
	defer db.Close()

	// Initialize with subdirectory packages
	equipmentWriter := equipment.NewWriter(db)
	equipmentReader := equipment.NewReader(db)

	attachmentWriter := attachment.NewWriter(db)
	attachmentReader := attachment.NewReader(db)

	muscleWriter := muscle.NewWriter(db)
	muscleReader := muscle.NewReader(db)

	muscleGroupWriter := musclegroup.NewWriter(db)
	muscleGroupReader := musclegroup.NewReader(db)

	categoryWriter := category.NewWriter(db)
	categoryReader := category.NewReader(db)

	subcategoryWriter := subcategory.NewWriter(db)
	subcategoryReader := subcategory.NewReader(db)

	app := application.New(
		equipmentWriter, equipmentReader,
		attachmentWriter, attachmentReader,
		muscleGroupWriter, muscleGroupReader,
		muscleWriter, muscleReader,
		categoryWriter, categoryReader,
		subcategoryWriter, subcategoryReader,
	)

	httpServer := web.NewApp(app, web.WithPort(8080))
	httpServer.Run()
}
