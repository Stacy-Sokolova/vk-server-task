package app

import (
	"context"
	"errors"
	"net/http"
	"vk-server-task/internal/metrics"
)

type Server struct {
	server       *http.Server
	metricServer *http.Server
	closer       *Closer
}

func NewServer(port string, handler http.Handler) *Server {
	s := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	metricMux := http.NewServeMux()
	metricMux.Handle("/metrics", metrics.MetricsHandler())
	metricServer := &http.Server{
		Addr:    ":9090",
		Handler: metricMux,
	}

	return &Server{
		server:       s,
		metricServer: metricServer,
		closer:       NewCloser(),
	}
}

func (s *Server) Run() error {
	s.closer.Add(func(ctx context.Context) error {
		//s.logger.Infow("Shutting down HTTP server")
		return s.server.Shutdown(ctx)
	})

	s.closer.Add(func(ctx context.Context) error {
		//s.logger.Infow("Shutting down Metrics server")
		return s.metricServer.Shutdown(ctx)
	})

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			//s.logger.Fatalw("HTTP server error",
			//	"error", err)
		}
	}()

	go func() {
		//s.logger.Infow("Starting Metrics server", "address", s.metricServer.Addr)
		if err := s.metricServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			//s.logger.Fatalw("Metrics server error",
			//	"error", err)
		}
	}()

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.closer.Close(ctx)
}
