package controllers

import (
	animeU "bootCampApi/api/usecases"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/structs"
	"github.com/gorilla/mux"
)

func GetAnimeByIdC(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	animeData := animeU.GetAnimeById(id)
	animeValues := make([]string, 0)
	for _, v := range structs.Values(animeData) {
		temp := fmt.Sprint(v)
		animeValues = append(animeValues, temp)
	}
	// open csv
	f, err := os.OpenFile("test.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Unable to open test.csv", err)
		responseHandle(w, nil, http.StatusInternalServerError)
		return
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println("Unable to close test.csv", err)
		}
	}(f)

	csvwriter := csv.NewWriter(f)
	fmt.Println(f)
	fmt.Println(csvwriter)
	defer csvwriter.Flush()

	if err := csvwriter.Write(animeValues); err != nil {
		log.Fatalln("error writing record to file", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ok...")
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ok...")
}
