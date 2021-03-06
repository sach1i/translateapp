package translateapp

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type App struct {
	Router  *mux.Router
	Service ServiceInterface
	Logger  *zap.SugaredLogger
}

type AppInterface interface {
	TranslateInput() http.HandlerFunc
	ListLanguages() http.HandlerFunc
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func NewApp(service ServiceInterface, logger *zap.SugaredLogger) *App {
	a := App{
		Router:  mux.NewRouter(),
		Service: service,
		Logger:  logger,
	}
	a.routes()
	return &a
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Router.ServeHTTP(w, r)
}
