package common

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"gobootcamp/models"
)

func CsvToPokemon(f multipart.File) (models.Pokemons, error) {
	var pokemons models.Pokemons

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return models.Pokemons{}, err
	}

	for _, item := range lines {
		id, err := strconv.Atoi(item[0])

		if err != nil {
			fmt.Println(err)
			continue
		}

		pokemon := models.Pokemon{
			Id:   id,
			Name: item[1],
		}

		pokemons = append(pokemons, pokemon)
	}

	fmt.Println(pokemons)
	return pokemons, nil
}

func HandleInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Some Error Occurred"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened. Err: %s", err)
	}
	w.Write(jsonResp)
}
