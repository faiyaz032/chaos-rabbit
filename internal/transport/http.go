package transport

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/faiyaz032/chaos-rabbit/internal/chaos"
	"github.com/faiyaz032/chaos-rabbit/internal/config"
)

type HTTPTransport struct {
	config config.HTTPConfig
}

func NewHTTPTransport(cfg config.HTTPConfig) *HTTPTransport {
	return &HTTPTransport{
		config: cfg,
	}
}

func (h *HTTPTransport) Start(ctx context.Context) error {
	proxy, err := createProxy(h.config.Target)
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr: h.config.Listen,
		// the latency will come from config
		Handler: chaos.Chain(proxy, chaos.Latency(2*time.Second)),
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
