package routes

import (
	"academy-go-q42021/controllers"

	"github.com/gorilla/mux"
)

//Get handler routes
func Get(router *mux.Router) {
	router.HandleFunc("/pokemons", controllers.GetAllPokemons)
	router.HandleFunc("/pokemons/{id}", controllers.GetPokemonById)
}
