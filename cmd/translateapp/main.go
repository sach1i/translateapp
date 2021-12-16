package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"translateapp/internal/logger"
	"translateapp/internal/server"
)

func main() {
	myServer := server.NewServer(logger.NewLogger()).Run()
	// channel for listening to myServer errors
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Server is listening on %s", myServer.Addr)
		serverErrors <- myServer.ListenAndServe()
	}()
	// channel for listening to system errors
	osShutdown := make(chan os.Signal, 1)
	signal.Notify(osShutdown, os.Interrupt, syscall.SIGTERM)
	// graceful shutdown differ based on source of error
	select {
	case err := <-serverErrors:
		log.Printf("myServer error: %s", err)

	case <-osShutdown:
		log.Printf("Server starting shutdown")
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		err := myServer.Shutdown(ctx)
		if err != nil {
			log.Printf("Not graceful shutdown, error: %v", err)
			_ = myServer.Close()
		}
		log.Printf("Server shutdown gracefully")
	}
}
