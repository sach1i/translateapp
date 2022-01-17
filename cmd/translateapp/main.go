package main

import (
	"go.uber.org/zap"
	"log"
	"time"
	"translateapp/internal/cache"
	"translateapp/internal/libretranslate"
	"translateapp/internal/logging"
	"translateapp/internal/server"
	"translateapp/internal/translateapp"
	"translateapp/internal/translator"
)

const BaseURLLibre = "http://libretranslate:5000"

func main() {
	newLogger := logging.NewLogger("INFO", true)
	if err := runApp(newLogger); err != nil {
		log.Fatal(err)
	}
}

func runApp(logger *zap.SugaredLogger) error {
	const (
		cacheSize = 20
		cacheTTL  = 10 * time.Second
	)

	client := libretranslate.NewClient(logger, BaseURLLibre)

	newCache := cache.NewCache(cacheSize, cacheTTL, logger)
	defer newCache.Close()

	newTranslator := translator.NewTranslator(newCache, client, logger)

	logic := translateapp.NewService(client, newTranslator, logger)

	app := translateapp.NewApp(logic, logger)

	srv := server.NewServer(app, logger)

	return server.RunServer(srv, logger)
}
