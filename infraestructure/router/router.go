package router

import (
	controller "mainRoot/interface/controller"
	"net/http"
)

const SearchPokemon = "/buscarPokemon"
const ListPokemon = "/listarPokemon"
const ListPokemonConcurrently = "/listarPokemonConcurrently"
const Port = ":8081"

func Listen() {
	http.HandleFunc(ListPokemon, func(w http.ResponseWriter, peticion *http.Request) {
		controller.ListarPokemones(w, peticion)
	})
	http.HandleFunc(SearchPokemon, func(w http.ResponseWriter, peticion *http.Request) {
		controller.BuscarPokemones(w, peticion)
	})

	http.HandleFunc(ListPokemonConcurrently, func(w http.ResponseWriter, peticion *http.Request) {
		controller.ListarPokemonesConcurrently(w, peticion)
	})

}
