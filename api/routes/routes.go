package routes

import (
	"fmt"
	"github.com/AndresCravioto/academy-go-q42021/api/services"
	"github.com/AndresCravioto/academy-go-q42021/controller"
	"github.com/AndresCravioto/academy-go-q42021/repositories"
	"github.com/AndresCravioto/academy-go-q42021/workerPool"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Routes() {
	searchService := services.NewSearchService(repositories.CreateAllChampionsList())
	apiService := services.NewWriteService(repositories.NewChampionsWriter())
	worker := workerPool.NewChampionWorker()
	apiController := controller.NewChampionsHandler(searchService, apiService, worker)
	router := mux.NewRouter()
	log.Println("champions api")
	api := router.PathPrefix("/championsApi/v1").Subrouter()
	api.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "championsApi v1")
	})

	api.HandleFunc("/champions/", apiController.ChampionsInformation).Methods(http.MethodGet)
	api.HandleFunc("/champions/{championId}", apiController.Champion).Methods(http.MethodGet)
	api.HandleFunc("/createChampionsDB/", apiController.DDragonChampionsList).Methods(http.MethodGet)
	api.HandleFunc("/worker/{type}/{items}/{items_per_worker}", apiController.ChampionsByWorker).Methods(http.MethodGet)

	log.Println(http.ListenAndServe(":8080", router))
}
