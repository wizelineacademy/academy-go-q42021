package services

import (
	"academy-go-q42021/models"
)

// all functions in this file should read from csv, but didn't achieved it atm

func GetAllPokemons() ([]models.Pokemon, error) {
	pikachu := models.Pokemon{1, "Pikachu"}
	chikorita := models.Pokemon{2, "Chikorita"}
	result := []models.Pokemon{pikachu, chikorita}
	return result, nil
}

func GetPokemonById(id string) (models.Pokemon, error) {
	charizard := models.Pokemon{2, "Charizard"}
	return charizard, nil
}
