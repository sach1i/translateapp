package main

import (
	"go.uber.org/zap"
	"log"
	"translateapp/internal/libretranslate"
	"translateapp/internal/logging"
	"translateapp/internal/server"
	"translateapp/internal/translateapp"
)

func main() {
	newLogger := logging.NewLogger("INFO", true)
	if err := runApp(newLogger); err != nil {
		log.Fatal(err)
	}
}

func runApp(logger *zap.SugaredLogger) error {
	client := libretranslate.NewClient(logger)

	logic := translateapp.NewService(client, logger)

	app := translateapp.NewApp(logic, logger)

	srv := server.NewServer(app, logger)
	if err := server.RunServer(srv, logger); err != nil {
		return err
	}
	return nil
}
