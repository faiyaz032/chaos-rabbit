package main

import (
	"context"
	"log"

	"github.com/faiyaz032/chaos-rabbit/internal/config"
	"github.com/faiyaz032/chaos-rabbit/internal/engine"
	"github.com/faiyaz032/chaos-rabbit/internal/transport"
)

func main() {
	cfg, err := config.Load("http.yaml")

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	e := engine.New()

	httpTransport := transport.NewHTTPTransport(cfg.HTTP)
	e.Register(httpTransport)
	e.Run(ctx)
}
