package routes

import (
	"fmt"
	"github.com/AndresCravioto/academy-go-q42021/api/services"
	"github.com/AndresCravioto/academy-go-q42021/controller"
	"github.com/AndresCravioto/academy-go-q42021/repositories"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Routes() {
	searchService := services.NewSearchService(repositories.CreateAllChampionsList())
	apiService := services.NewWriteService(repositories.NewChampionsWriter())
	apiController := controller.NewChampionsHandler(searchService, apiService)
	router := mux.NewRouter()
	log.Println("champions api")
	api := router.PathPrefix("/championsApi/v1").Subrouter()
	api.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "championsApi v1")
	})

	api.HandleFunc("/champions/", apiController.ChampionsInformation).Methods(http.MethodGet)
	api.HandleFunc("/champions/{championId}", apiController.Champion).Methods(http.MethodGet)
	api.HandleFunc("/addChampion/", apiController.ChampionList).Methods(http.MethodGet)
	log.Println(http.ListenAndServe(":8081", router))
}
