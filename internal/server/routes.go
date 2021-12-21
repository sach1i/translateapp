package server

func (s *Server) routes() {
	s.HandleFunc("/translate", s.TranslateInput()).Methods("POST")
	s.HandleFunc("/languages", s.ListLanguages()).Methods("GET")
}
