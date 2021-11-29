package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Input struct{
	Word string `json:"word"`
	Source string `json:"source"`
	Target string `json:"target"`
}
type Response struct {
	TranslatedWord string
}


type Server struct {
	*mux.Router

}

func NewServer() *Server{
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes(){
	s.HandleFunc("/translate",s.translateInput()).Methods("POST")
	s.HandleFunc("/languages",s.listLanguages()).Methods("GET")
}

func (s *Server) translateInput() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var i Input
		if err := json.NewDecoder(r.Body).Decode(&i);err!=nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp := Response{}
		resp.TranslatedWord = i.Word
		respJson, err := json.Marshal(resp); if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type","application/json")
		w.Write(respJson)
	}
}

func (s *Server) listLanguages() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		resp := make(map[string]string)
		resp["code"] = "code"
		resp["name"] = "name"
		respJson, err := json.Marshal(resp); if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(respJson)
	}
}
