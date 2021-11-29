package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"translateapp/internal/app/apiserver"
)

func main() {
	log.Printf("starting...")
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	srv:= apiserver.NewServer()
	http.ListenAndServe(":8080",srv)
	defer done()
	<-ctx.Done()
	log.Printf("successful shutdown")


}