// Package db
package db

import (
	"github.com/cheezecakee/fitrkr/atlas/internal/core/ports"
)

type PortsProvider interface {
	CreatePorts() (ports.Read, ports.Write)
}
