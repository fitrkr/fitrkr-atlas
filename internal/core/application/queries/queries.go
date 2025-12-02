// Package queries provides all command handlers for the application.
//
// Queries are automatically registered when the package is imported.
// Each query file contains an init() function that registers itself.
//
// Usage:
//  1. Import the package
//  2. Call Init(write, read) to set dependencies
//  3. Call RegisterAll(registry) to register all queries
//
// To add a new query, create a file with:
//   - Query struct with Handle method
//   - init() function calling register()
package queries

import (
	"github.com/fitrkr/atlas/internal/core/application/mediator"
	"github.com/fitrkr/atlas/internal/core/ports"
)

var (
	read     ports.Read
	handlers []mediator.Handler
)

func Init(r ports.Read) {
	read = r
}

func RegisterAll(reg *mediator.Registry) {
	for _, h := range handlers {
		reg.Register(h)
	}
}

func register(h mediator.Handler) {
	handlers = append(handlers, h)
}
