package mediator

import (
	"context"
	"reflect"
)

type Registry struct {
	handle map[reflect.Type]Handler
}

func NewRegistry() *Registry {
	return &Registry{
		handle: make(map[reflect.Type]Handler),
	}
}

func (r *Registry) Register(msg Handler) {
	r.handle[reflect.TypeOf(msg)] = msg
}

func (r *Registry) Send(ctx context.Context, msg Handler) (any, error) {
	return msg.Handle(ctx)
}
