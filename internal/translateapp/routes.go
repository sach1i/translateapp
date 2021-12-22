package translateapp

func (a *App) routes() {
	a.Router.HandleFunc("/translate", a.TranslateInput()).Methods("POST")
	a.Router.HandleFunc("/languages", a.ListLanguages()).Methods("GET")
}
