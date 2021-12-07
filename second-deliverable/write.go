package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	entity "academy-go-q42021/second-deliverable/pkg/entity"
)

// Items - Description here....
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func writeCSV(str string) error {
	//str := `[{"id": 1,"name": "Pancho Lopez"},{"id": 2,"name": "Juana Juarez"},{"id": 4,"name": "Lola Contreras"},{"id": 6,"name": "Tomasa Perez"}]`

	var jsonItemsList []Item

	if err := json.Unmarshal([]byte(str), &jsonItemsList); err != nil {
		panic(err)
	}
	fmt.Println(jsonItemsList)

	outputFile, err := os.Create("output.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	for _, r := range jsonItemsList {
		var csvRow []string
		csvRow = append(csvRow, fmt.Sprint(r.ID), r.Name)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func writeCsvData(data []byte) error {

	var r entity.JSONOutput
	err := json.Unmarshal(data, &r)

	outputFile, err := os.Create("output.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	for _, pokemon := range r.PokemonEntries {
		var csvRow []string
		csvRow = append(csvRow, fmt.Sprint(pokemon.EntryNumber), pokemon.PokemonSpecies.Name, pokemon.PokemonSpecies.URL)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
		fmt.Printf("%d%s%s%s%s\n", pokemon.EntryNumber, ",", pokemon.PokemonSpecies.Name, ",", pokemon.PokemonSpecies.URL)
	}

	return nil
}

func UnmarshalWelcome(data []byte) (entity.JSONOutput, error) {
	var r entity.JSONOutput
	err := json.Unmarshal(data, &r)
	fmt.Println(r.Name)
	fmt.Println(len(r.PokemonEntries))
	for _, pokemon := range r.PokemonEntries {
		fmt.Printf("%d%s%s%s%s\n", pokemon.EntryNumber, ",", pokemon.PokemonSpecies.Name, ",", pokemon.PokemonSpecies.URL)
	}
	return r, err
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")
	//response, err := http.Get("http://localhost:8080/items")
	if err != nil {
		fmt.Printf("HTTP Fail with error %s\n", err)

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		writeCsvData(data)
	}

}
