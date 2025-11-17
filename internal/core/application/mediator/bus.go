// Package mediator
package mediator

import (
	"context"
)

type Bus interface {
	Send(ctx context.Context, msg Handler) (any, error)
}

type Handler interface {
	Handle(ctx context.Context) (any, error)
}
