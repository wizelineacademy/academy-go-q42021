package main

import (
	entity "academy-go-q42021/final-deliverable/pkg/entity"
	repo "academy-go-q42021/final-deliverable/pkg/repo"
	server "academy-go-q42021/final-deliverable/pkg/server"
	workerpool "academy-go-q42021/final-deliverable/pkg/workerpool"
	"sync"
	"time"

	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var m = sync.Mutex{}

// Description: * An endpoint for reading from an external API
func createCSVFile() {
	fmt.Println(">> Create a CSV File getting the entries from an external endpoint")
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Printf("HTTP Fail with error %s\n", err)

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		writeCSVData(data)
	}
}

// Description: * Write the information in a CSV file
func writeCSVData(data []byte) error {
	fmt.Println(">> Write the entries to a CSV File")
	var r entity.JSONOutput
	err := json.Unmarshal(data, &r)

	outputFile, err := os.Create("output.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	for _, entry := range r.PokemonEntries {
		var csvRow []string
		csvRow = append(csvRow, fmt.Sprint(entry.EntryNumber), entry.PokemonSpecies.Name, entry.PokemonSpecies.URL)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

// Description * An endpoint for reading the CSV
func readCSVFile() []entity.Item {
	fmt.Println(">> Read the entries to a CSV File")
	csvFile, err := os.Open("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var itemsList []entity.Item

	for _, line := range csvLines {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}
		record := entity.Item{
			ID: id, Name: line[1], URL: line[2],
		}
		itemsList = append(itemsList, record)
	}
	return itemsList
}

// read all the items form the repository
func getAllItems() {
	fmt.Println(">> Get the all entries stored in the repository and print as JSON")
	response, err := http.Get("http://localhost:8080/items")
	if err != nil {
		fmt.Printf("HTTP Fail with error %s\n", err)

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

// Read the CSV list using the worker pool strategy
func readCSVConcurrency(csvList []entity.Item) {
	fmt.Println(">> Read the CSV list using a worker pool")
	var allTask []*workerpool.Task
	cursor := 0

	for i := 1; i <= 100; i++ {
		task := workerpool.NewTask(
			func(data interface{}) error {
				taskID := data.(int)
				time.Sleep(100 * time.Millisecond)
				m.Lock()
				s := csvList[cursor]
				fmt.Printf("Task %d processed, string:%v \n", taskID, s)
				cursor++
				m.Unlock()
				return nil
			}, i)
		allTask = append(allTask, task)
	}

	pool := workerpool.NewPool(allTask, 5)
	pool.Run()
}

// Run the all the functions of the program
func run() (int, error) {

	fmt.Println(">>> Start program...")

	createCSVFile() // write a CSV file

	csvList := readCSVFile() // read CSV file

	repository := repo.NewItemsRepository(csvList)
	s := server.New(repository)
	go func() {

		fmt.Println("The items server is on tap now: http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", s.Router()))
	}()

	getAllItems() //An endpoint for reading the CSV

	readCSVConcurrency(csvList) //An endpoint for reading the CSV concurrently
	fmt.Println(">>> End of program ... ")
	return 0, nil
}

// The main code starts here
func main() {
	if rc, err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(rc)
	}
}
