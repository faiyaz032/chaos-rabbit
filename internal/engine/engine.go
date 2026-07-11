package engine

import (
	"context"

	"github.com/faiyaz032/chaos-rabbit/internal/transport"
)

type Engine struct {
	transports []transport.Transport
}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Register(t transport.Transport) {
	e.transports = append(e.transports, t)
}

func (e *Engine) Run(ctx context.Context) error {
	for _, t := range e.transports {
		go func(t transport.Transport) {
			t.Start(ctx)
		}(t)
	}
	<-ctx.Done()
	return nil
}
