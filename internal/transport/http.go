package transport

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/faiyaz032/chaos-rabbit/internal/chaos"
	"github.com/faiyaz032/chaos-rabbit/internal/config"
)

type HTTPTransport struct {
	config config.HTTPConfig
	chaos  []config.ChaosConfig
}

func NewHTTPTransport(cfg config.HTTPConfig, chaosCfg []config.ChaosConfig) *HTTPTransport {
	return &HTTPTransport{
		config: cfg,
		chaos:  chaosCfg,
	}
}

func (h *HTTPTransport) Start(ctx context.Context) error {
	proxy, err := createProxy(h.config.Target)
	if err != nil {
		return err
	}

	middlewares, err := chaos.Build(h.chaos)

	server := &http.Server{
		Addr: h.config.Listen,
		// the latency will come from config
		Handler: chaos.Chain(proxy, middlewares...),
	}

	go func() {
		<-ctx.Done()
		server.Shutdown(context.Background())
	}()

	return server.ListenAndServe()
}

func createProxy(target string) (*httputil.ReverseProxy, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(targetURL), nil
}
