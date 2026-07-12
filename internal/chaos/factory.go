package chaos

import (
	"fmt"
	"time"

	"github.com/faiyaz032/chaos-rabbit/internal/config"
)

func Build(specs []config.ChaosConfig) ([]Middleware, error) {
	middlewares := make([]Middleware, 0, len(specs))
	for _, spec := range specs {
		mw, err := build(spec)
		if err != nil {
			return nil, err
		}
		middlewares = append(middlewares, mw)
	}
	return middlewares, nil
}

func build(spec config.ChaosConfig) (Middleware, error) {
	switch spec.Type {
	case "latency":
		return buildLatency(spec.Config)
	default:
		return nil, fmt.Errorf("chaos: unknown type %q", spec.Type)
	}
}

func buildLatency(cfg map[string]any) (Middleware, error) {
	raw, ok := cfg["duration"]
	if !ok {
		return nil, fmt.Errorf("chaos: latency requires 'duration'")
	}
	str, ok := raw.(string)
	if !ok {
		return nil, fmt.Errorf("chaos: latency 'duration' must be a string (e.g. \"2s\")")
	}
	d, err := time.ParseDuration(str)
	if err != nil {
		return nil, fmt.Errorf("chaos: invalid latency duration: %w", err)
	}
	return Latency(d), nil
}
