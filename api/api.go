package api

import (
	controllers "bootCampApi/api/controllers"

	"github.com/gorilla/mux"
)

func New() *mux.Router {

	router := mux.NewRouter()

	routes(router)
	return router
}

func routes(router *mux.Router) {
	router.HandleFunc("/", controllers.HealthCheck).Methods("GET")
	router.HandleFunc("/readCSV", controllers.ReadCSV).Methods("GET")
	router.HandleFunc("/readCSV/{id:[0-9]+}", controllers.ReadCSV).Methods("GET")
	router.HandleFunc("/writeCSV/{id:[0-9]+}", controllers.GetAnimeByIdC).Methods("GET")
}
