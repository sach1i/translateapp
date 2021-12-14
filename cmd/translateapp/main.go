package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"translateapp/internal/server"
)

func main() {
	server := server.NewServer().Run()
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
