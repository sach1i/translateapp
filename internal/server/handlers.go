package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"translateapp/internal/models"
	"translateapp/internal/service"
)

// TranslateInput takes request payload and returns translated word
func (s *Server) TranslateInput() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		myService := service.NewService()
		ctx := context.Background()
		w.Header().Set("Content-Type", "application/json")
		var i models.Input
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			log.Fatal(err)
		}

		clientResponse, customErr := myService.Translate(ctx, i)
		if customErr != nil {
			err := json.NewEncoder(w).Encode(customErr)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := json.NewEncoder(w).Encode(clientResponse)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Printf("exec time: %s", time.Now().Sub(startTime))
	}
}

// ListLanguages returns list of all available languages for translation in form {"code":*,"name":*}
func (s *Server) ListLanguages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		startTime := time.Now()
		myService := service.NewService()
		w.Header().Set("Content-Type", "application/json")
		log.Printf("GET request triggered at %s", startTime)
		clientResponse, customErr := myService.Languages(ctx)
		if customErr != nil {
			err := json.NewEncoder(w).Encode(customErr)
			if err != nil {
				log.Fatal(err)
			}

		} else {
			err := json.NewEncoder(w).Encode(clientResponse)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Printf("exec time: %s", time.Now().Sub(startTime))
	}
}
