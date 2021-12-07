package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Input struct {
	Word   string `json:"word"`
	Source string `json:"source"`
	Target string `json:"target"`
}
type Response struct {
	TranslatedWord string
}

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/translate", s.TranslateInput()).Methods("POST")
	s.HandleFunc("/languages", s.ListLanguages()).Methods("GET")
}

// TranslateInput takes request payload and returns translated word
func (s *Server) TranslateInput() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Post request triggered at %s", startTime)
		var i Input
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var resp Response

		resp.TranslatedWord = i.Word

		respJson, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJson)
		log.Printf("User input was: %v\nexec time: %s", i, time.Now().Sub(startTime))
	}
}

// ListLanguages returns list of all available languages for translation in form {"code":*,"name":*}
func (s *Server) ListLanguages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("GET request triggered at %s", startTime)
		w.Header().Set("Content-Type", "application/json")
		var resp = make(map[string]string)
		resp["code"] = "code"
		resp["name"] = "name"
		respJson, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(respJson)
		log.Printf("exec time: %s", time.Now().Sub(startTime))
	}
}
