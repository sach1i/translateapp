package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	log.Printf("starting...")
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer done()
	<-ctx.Done()
	log.Printf("successful shutdown")


}