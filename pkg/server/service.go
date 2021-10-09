package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Service is the server abstraction as a service.
type Service interface {
	ListenAndServe(duration time.Duration) error
}

type httpService struct {
	server *http.Server
}

// ListenAndServe runs the server and handles incoming requests. The server can
// be manually shutdown with system calls like: interrupt, termination or kill signals.
// The server is then gracefully shutdown with the given timeout. After the timeout
// exceeds the server is forcefully shutdown.
func (s *httpService) ListenAndServe(timeout time.Duration) (err error) {
	go func() {
		// non blocking server
		err = s.server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
	}()

	quit := make(chan os.Signal, 1)
	// SIGKILL cannot be handled
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// blocking
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server is shutting down: %w", err)
	}

	return nil
}

// NewService returns a Service implementation with the given HTTP server.
func NewService(server *http.Server) Service {
	return &httpService{
		server: server,
	}
}