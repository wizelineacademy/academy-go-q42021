package routes

import (
	"fmt"
	"github.com/AndresCravioto/academy-go-q42021/api/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Routes() {
	router := mux.NewRouter()
	log.Println("stands api")
	api := router.PathPrefix("/standsApi/v1").Subrouter()
	api.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "standsApi v1")
	})

	api.HandleFunc("/stands/", services.GetAllStands).Methods(http.MethodGet)
	api.HandleFunc("/stands/{standId}", services.SearchByStandId).Methods(http.MethodGet)
	log.Println(http.ListenAndServe(":8081", router))
}
