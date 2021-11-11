package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"academy-go-q42021/services"

	"github.com/gorilla/mux"
)

func GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	pokemons, err := services.GetAllPokemons()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(pokemons)
}

func GetPokemonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	pokemons, err := services.GetPokemonById(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(pokemons)

}
