package controllers

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	animeI "bootCampApi/api/interfaces"

	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ok...")
}

func ReadCSV(w http.ResponseWriter, r *http.Request) {
	// check params
	id := mux.Vars(r)["id"]
	idValue := 0
	if id != "" {
		row, err := strconv.Atoi(id)
		if err != nil {
			return
		}
		idValue = row
	}

	// open csv
	f, err := os.Open("test.csv")
	if err != nil {
		log.Println("Unable to read test.csv", err)
		responseHandle(w, nil, http.StatusInternalServerError)
		return
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println("Unable to close test.csv", err)
		}
	}(f)

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		responseHandle(w, nil, http.StatusInternalServerError)
		return
	}
	var newRecord []string
	for i := 0; i < len(records); i++ {
		value, _ := strconv.Atoi(records[i][0])
		if value == idValue {
			newRecord = records[i]
			break
		}
	}

	if id != "" && len(newRecord) == 0 {
		log.Println("Record does not exists")
		responseHandle(w, newRecord, http.StatusAccepted)
		return
	}
	if len(newRecord) > 1 {
		x := [][]string{newRecord}
		records = x
	}

	var response = make([]animeI.AnimeStruct, len(records))
	for row, content := range records {

		animeId, err := strconv.Atoi(content[0])
		if err != nil {
			responseHandle(w, nil, http.StatusInternalServerError)
			return
		}

		singleRow := animeI.AnimeStruct{
			AnimeId:  animeId,
			Title:    content[1],
			Synopsis: content[2],
			Studio:   content[3],
		}

		response[row] = singleRow
	}

	responseHandle(w, response, http.StatusOK)
}
