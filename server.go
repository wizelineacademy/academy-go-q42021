package main

import (
	"fmt"
	"github.com/AndresCravioto/academy-go-q42021/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var (
	stands services.StandServices
)


func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	stands = &services.Stands{}
	stands.Initialize()
}

func main() {
	r := mux.NewRouter()
	log.Println("stands api")
	api := r.PathPrefix("/standsApi/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "standsApi v1")
	})

	api.HandleFunc("/stands/id/{standId}", searchByStandId).Methods(http.MethodGet)
	log.Fatalln(http.ListenAndServe(":8080", r))
}