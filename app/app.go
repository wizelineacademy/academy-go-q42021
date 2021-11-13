package app

import (
	"GOBootcamp/app/database"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	CSV    database.MyCSV
}

func New() *App {

	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/getContentCSV", a.GetPostHandler()).Methods("GET")
}
