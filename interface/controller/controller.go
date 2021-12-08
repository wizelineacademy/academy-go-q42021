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
	for _, v := range pokedex {
		pokeJson, err2 := pokemonToJson(v)
		if err2 {
			fmt.Println("Error ")
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, pokeJson)
	}
}

func BuscarPokemones(w http.ResponseWriter, peticion *http.Request) {
	id, existeId := peticion.URL.Query()["id"]
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
		}
	} else {
		io.WriteString(w, "Id no encontrado")
	}
}

func buscarPokemonPorId(id int) (Model.Pokemon, bool) {
	for _, v := range pokedex {
		if v.Id == id {
			return v, false
		}
	}
	poke := Model.Pokemon{-1, "", ""}
	return poke, true
}

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
		id, err := strconv.Atoi(row[0])
		if err == nil {

		}
		poke := Model.Pokemon{id, row[1], row[2]}
		pokedex = append(pokedex, poke)

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

//Deliverable 3 FINAL! (STARTS HERE)
func ListarPokemonesConcurrently(w http.ResponseWriter, peticion *http.Request) {
	err, items, evenOdd, itemsPerWorker := ListarPokemonesConcurrently_VerifyParams(peticion)
	if err != "" {
		fmt.Println("Error: " + err)
		io.WriteString(w, err)
		return
	}
	newPokedex := ReadCsvConcurrently(items, evenOdd, itemsPerWorker)

	ListarPokemonesEnviados(w, peticion, newPokedex)
}

func ReadCsvConcurrently(items int, evenOddType string, itemsPerWorker int) []Model.Pokemon {
	var newPokedex []Model.Pokemon
	newPokedex = []Model.Pokemon{}
	var evenOdd int
	if evenOddType == "even" {
		evenOdd = 0
	} else {
		evenOdd = 1
	}
	// Listo para leer el archivo
	fileName := "pokemon.csv"
	extension := strings.Split(fileName, ".")
	if extension[1] != "csv" {
		log.Fatal("Archivo no soportado")
		return nil
	}

	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
		return nil
	}

	defer fs.Close()
	//Nueva lógica
	var contWorkers = (items / itemsPerWorker) + 1
	jobs := make(chan int, items)
	results := make(chan Model.Pokemon, items)

	for i := 0; i < contWorkers; i++ {
		go func(jobs <-chan int, results chan<- Model.Pokemon) {
			for j := range jobs {
				poke, err := buscarPokemonPorId(parN(evenOdd, j)) //buscamos el pokemon por su id utilizando n numero (par o impar), siendo 2 par el número 4. Y 2 impar el número 3
				if err {
					fmt.Println(parN(evenOdd, j))
					fmt.Println("Id no encontrado")
				}
				results <- poke
			}
		}(jobs, results)
	}

	for j := 0; j < items; j++ {
		jobs <- j
	}
	close(jobs)
	for j := 0; j < items; j++ {
		newPokedex = append(newPokedex, <-results)
	}
	return newPokedex
}

func parN(evenOdd int, n int) int {
	if evenOdd == 0 {
		return 2 * n
	}
	if evenOdd == 1 {
		if n == 0 {
			return 1
		}
		return (2 * n) - 1
	}
	return 0
}
func ListarPokemonesEnviados(w http.ResponseWriter, peticion *http.Request, pokes []Model.Pokemon) {
	for _, v := range pokes {
		pokeJson, err2 := pokemonToJson(v)
		if err2 {
			fmt.Println("Error ")
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, pokeJson)
	}
}

func getRow(r *csv.Reader, evenOdd int, maxId int) (string, Model.Pokemon) {
	row, err := r.Read()

	if err != nil && err != io.EOF {
		return "can not  read, er r is " + err.Error(), Model.Pokemon{}
	}

	if err == io.EOF {
		return "EOF", Model.Pokemon{}
	}
	id, err := strconv.Atoi(row[0])
	if err != nil {
		return "Not Valid ID", Model.Pokemon{}
	}
	fmt.Println(row[1])
	if id > maxId {
		return "ID too high", Model.Pokemon{}
	}

	if id%2 == evenOdd { //si evenOdd es 1, entones buscamos impares, si el residuo es 1 es impar y prosigue, lo mismo aplica para par pero con 0
		poke := Model.Pokemon{id, row[1], row[2]}
		return "", poke
	} else {
		return getRow(r, evenOdd, maxId)
	}
}

//verificación de parámetros
//return error, items,
func ListarPokemonesConcurrently_VerifyParams(peticion *http.Request) (string, int, string, int) {
	itemsParam, itemExists := peticion.URL.Query()["items"]
	evenOddParam, typeExists := peticion.URL.Query()["type"]
	itemsPerWorkerParam, itemsPerWorkerExists := peticion.URL.Query()["items_per_worker"]

	//validación de  parámetro Items
	if !itemExists { //validar si existe el parámetro
		return "Parameter 'items' is required", 0, "", 0
	}
	// si existe buscamos que el valor sea correcto
	items, err := strconv.Atoi(itemsParam[0])
	if err != nil { //el parámetro debe ser numérico.
		return "Parameter 'items' must be number.", 0, "", 0
	}
	if items <= 0 { //El parámetro debe ser mayor a 0
		return "Parameter 'items' must be higher than 0.", 0, "", 0
	}
	if items > len(pokedex) {
		return "Parameter 'items' must be lower than " + strconv.Itoa(len(pokedex)) + " .", 0, "", 0
	}
	//Validación del parámetro Type
	if !typeExists { //Validar si existe el parámetro Type
		return "Parameter 'type' is required", 0, "", 0
	}
	// si existe, buscamos que el valor sea Odd o Even, únicamente.
	if strings.ToLower(evenOddParam[0]) != strings.ToLower("odd") && strings.ToLower(evenOddParam[0]) != strings.ToLower("even") {
		return "Parameter 'type' must be 'Even' or 'Odd'.", 0, "", 0
	}

	//Validación de parámetro 'items_per_worker'
	if !itemsPerWorkerExists { //validar si existe el parámetro 'items_per_worker'
		return "Parameter 'items_per_worker' is required", 0, "", 0
	}
	// si existe, buscamos que sea un valr numérico mayor a 0.
	itemsPerWorker, err := strconv.Atoi(itemsPerWorkerParam[0])
	if err != nil { //el parámetro debe ser numérico.
		return "Parameter 'items_per_worker' must be number.", 0, "", 0
	}
	if itemsPerWorker < 0 { //El parámetro debe ser mayor a 0
		return "Parameter 'items_per_worker' must be higher than 0.", 0, "", 0
	}
	return "", items, strings.ToLower(evenOddParam[0]), itemsPerWorker
}
