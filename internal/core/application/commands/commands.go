// Package commands provides all command handlers for the application.
//
// Commands are automatically registered when the package is imported.
// Each command file contains an init() function that registers itself.
//
// Usage:
//  1. Import the package
//  2. Call Init(write, read) to set dependencies
//  3. Call RegisterAll(registry) to register all commands
//
// To add a new command, create a file with:
//   - Command struct with Handle method
//   - init() function calling register()
package commands

import (
	"github.com/cheezecakee/fitrkr/atlas/internal/core/application/mediator"
	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

var (
	write    ports.Write
	read     ports.Read
	handlers []mediator.Handler
)

func Init(w ports.Write, r ports.Read) {
	write = w
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
