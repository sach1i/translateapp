package main

import (
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
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
	var wg = sync.WaitGroup{}
	if err := runApp(newLogger, &wg); err != nil {
		log.Fatal(err)
	}
}

func runApp(logger *zap.SugaredLogger, wg *sync.WaitGroup) error {
	var cacheSize = 20
	var cacheTTL = 10
	osShutdown := make(chan os.Signal, 1)
	signal.Notify(osShutdown, os.Interrupt, syscall.SIGTERM)
	wg.Add(2)
	client := libretranslate.NewClient(logger, BaseURLLibre)
	newCache := cache.NewCache(cacheSize, logger)
	newCache.Refresher(cacheTTL, wg, osShutdown)
	newTranslator := translator.NewTranslator(newCache, client, logger)
	logic := translateapp.NewService(client, newTranslator, logger)
	app := translateapp.NewApp(logic, logger)
	srv := server.NewServer(app, logger)
	if err := server.RunServer(srv, logger, wg, osShutdown); err != nil {
		return err
	}
	wg.Wait()
	return nil
}
