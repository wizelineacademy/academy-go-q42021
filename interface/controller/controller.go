package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	Model "mainRoot/domain/model"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var pokedex []Model.Pokemon

func LlenarPokedex() {
	if len(pokedex) == 0 {
		ReadCsv()
		fmt.Printf("%v\n", "Pokedex guardada")
	} else {
		fmt.Printf("%v\n", "Pokedex ya habia sido guardada")
	}
}

func ListarPokemones(w http.ResponseWriter, peticion *http.Request) {
	for i, v := range pokedex {
		//fmt.Printf("2**%d = %d\n", i, v)
		if i != 0 {
			//io.WriteString(w, "{"+pokemonToString(v)+"}")
			pokeJson, err2 := pokemonToJson(v)
			if err2 {
				fmt.Println("Error ")
			}
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, pokeJson)
		}
	}
}

func BuscarPokemones(w http.ResponseWriter, peticion *http.Request) {
	id, existeId := peticion.URL.Query()["id"]
	// Imprimir para depurar
	fmt.Printf("%v\n", id[0])
	if existeId {
		i, err := strconv.Atoi(id[0])
		if err == nil {
			fmt.Println(i)
		}
		poke, err1 := buscarPokemonPorId(i)
		if err1 {
			io.WriteString(w, "Id no encontrado")
		} else {
			w.Header().Set("Content-Type", "application/json")
			pokeJson, err2 := pokemonToJson(poke)
			if err2 {
				fmt.Println("Error ")
			}
			io.WriteString(w, pokeJson)
			//io.WriteString(w, pokemonToString(poke)+" \n")
		}
	} else {
		io.WriteString(w, "Id no encontrado")
	}
}

func buscarPokemonPorId(id int) (Model.Pokemon, bool) {
	for _, v := range pokedex {
		//fmt.Printf("2**%d = %d\n", i, v)
		if v.Id == id {
			return v, false
		}
	}
	poke := Model.Pokemon{-1, "", ""}
	return poke, true
}

/*
func LeerCvs() map[int]Model.Pokemon {
	var Pokedex = map[int]Model.Pokemon{
		1: {1, "Bulbasaur", "Leaf"},
		2: {2, "Yvisaur", "Leaf"},
		3: {3, "Venasaur", "Leaf"},
	}
	return Pokedex
}*/

func ReadCsv() {
	// Listo para leer el archivo
	fileName := "pokemon.csv"
	extension := strings.Split(fileName, ".")
	if extension[1] != "csv" {
		log.Fatal("Archivo no soportado")
		return
	}

	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		//fmt.Println(row)
		id, err := strconv.Atoi(row[0])
		if err == nil {

		}
		poke := Model.Pokemon{id, row[1], row[2]}
		pokedex = append(pokedex, poke)

		fmt.Println(poke)
	}
}

func pokemonToString(poke Model.Pokemon) string {
	return "Pokedex ID: " + strconv.Itoa(poke.Id) + " Pokemon name: " + poke.Name + " Type: " + poke.PokemonType
}
func pokemonToJson(poke Model.Pokemon) (string, bool) {
	j, err := json.Marshal(poke)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		return string(j), false
	}
	return "", true
}
