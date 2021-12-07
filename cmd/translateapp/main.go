package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"translateapp/internal/app/apiserver"
)

func main() {
	server := http.Server{
		Addr:         ":8080",
		Handler:      apiserver.NewServer(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	// channel for listening to server errors
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Server is listening on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()
	// channel for listening to system errors
	osShutdown := make(chan os.Signal, 1)
	signal.Notify(osShutdown, os.Interrupt, syscall.SIGTERM)
	// graceful shutdown differ based on source of error
	select {
	case err := <-serverErrors:
		log.Printf("server error: %s", err)

	case <-osShutdown:
		log.Printf("Server starting shutdown")
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Printf("Not graceful shutdown, error: %v", err)
			_ = server.Close()
		}
		log.Printf("Server shutdown gracefully")
	}
}
