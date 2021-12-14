package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"translateapp/internal/service"
)

// TranslateInput takes request payload and returns translated word
/*
func (s *Server) TranslateInput() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		service := service.NewService()
		log.Printf("Post request triggered at %s", startTime)
		var i models.Input
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp := service.Translate(i)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
		log.Printf("User input was: %v\nexec time: %s", i, time.Now().Sub(startTime))
	}
}
*/
// ListLanguages returns list of all available languages for translation in form {"code":*,"name":*}
func (s *Server) ListLanguages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		startTime := time.Now()
		service := service.NewService()
		log.Printf("GET request triggered at %s", startTime)
		w.Header().Set("Content-Type", "application/json")
		clientResponse, errorResponse := service.Languages(ctx)
		if errorResponse != nil {
			resp, err := json.Marshal(errorResponse)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(resp)
		}

		resp, err := json.Marshal(clientResponse)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s, %v", resp, resp)
		w.Write(resp)

		log.Printf("exec time: %s", time.Now().Sub(startTime))

	}
}
