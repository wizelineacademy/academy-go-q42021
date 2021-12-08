package main

import (
	"fmt"
	"log"
	router "mainRoot/infraestructure/router"
	controller "mainRoot/interface/controller"
	"net/http"
)

//
func main() {
	controller.LlenarPokedex()

	http.HandleFunc(router.ListPokemon, func(w http.ResponseWriter, peticion *http.Request) {
		controller.ListarPokemones(w, peticion)
	})
	http.HandleFunc(router.SearchPokemon, func(w http.ResponseWriter, peticion *http.Request) {
		controller.BuscarPokemones(w, peticion)
	})

	http.HandleFunc(router.ListPokemonConcurrently, func(w http.ResponseWriter, peticion *http.Request) {
		controller.ListarPokemonesConcurrently(w, peticion)
	})

	fmt.Println("Servidor listo escuchando en " + router.Port)
	log.Fatal(http.ListenAndServe(router.Port, nil))
}
