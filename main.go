package main

import (
	"academy-go-q42021/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	routes.Get(router)
	http.ListenAndServe(":80", router)
}
