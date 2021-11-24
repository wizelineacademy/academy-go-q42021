package api

import (
	"github.com/gorilla/mux"
)

func New() *mux.Router {

	router := mux.NewRouter()

	routes(router)
	return router
}

func routes(router *mux.Router) {
	router.HandleFunc("/", HealthCheck).Methods("GET")
	router.HandleFunc("/readCSV", ReadCSV).Methods("GET")
	router.HandleFunc("/readCSV/{id:[0-9]+}", ReadCSV).Methods("GET")
}
