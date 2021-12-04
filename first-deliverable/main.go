package main

import (

	"fmt"
	"log"
	"os"	
	"strconv"	
	"encoding/csv"
	"net/http"
	repo "academy-go-q42021/pkg/repo"
	server "academy-go-q42021/pkg/server"
	entity "academy-go-q42021/pkg/entity"
)

// Items - Define type to create a repository for items
type Items []entity.Item

func readCSVFile() []entity.Item  {
		// open the CSV FILE and read and print it contain.
		csvFile, err := os.Open("dataFile.csv")
		if err != nil {
			log.Fatal( err )
		}
		fmt.Println("Successfully Opened CSV file")
		defer csvFile.Close()
	
		// read
		csvLines, err := csv.NewReader(csvFile).ReadAll()
		if err != nil {
			 log.Fatal( err )
		}
	
		var itemsList []entity.Item
		for _, line := range csvLines {
			id, err := strconv.Atoi(line[0])
	
			if err != nil {
				log.Fatal( err )
			}
			record := entity.Item{
				ID:   id,
				Name: line[1],
			}
			itemsList = append(itemsList, record)
		}	
		return itemsList
	}
	
func main() {
	if rc, err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(rc)
	}
}

func run() (int, error) {

	fmt.Println("running main")

	csvList := readCSVFile()		// read CSV file

	repository := repo.NewItemsRepository(csvList)

	s := server.New( repository )

	fmt.Println("The items server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))

	return 0, nil
}
