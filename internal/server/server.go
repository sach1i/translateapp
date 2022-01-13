package server

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func NewServer(app http.Handler, logger *zap.SugaredLogger) *http.Server {
	server := http.Server{
		Addr:         ":8080",
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	logger.Info("server starting")
	return &server
}

func RunServer(server *http.Server, logger *zap.SugaredLogger) error {
	serverErrors := make(chan error, 1)
	go func() {
		logger.Infof("server is listening at: %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()
	// channel for listening to system errors
	osShutdown := make(chan os.Signal, 1)
	signal.Notify(osShutdown, os.Interrupt, syscall.SIGTERM)
	// graceful shutdown differ based on source of error
	select {
	case err := <-serverErrors:
		logger.Fatalf("error starting server: %s", err)
		return fmt.Errorf("error starting server: %s", err)

	case <-osShutdown:
		logger.Info("Shutdown started")
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			err = server.Close()
		}
	}
	return nil
}
