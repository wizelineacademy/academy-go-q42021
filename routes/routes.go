package routes

import (
	"encoding/csv"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gorilla/mux"
)

type Model struct {
	Id   int
	Name string
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: mainHandler")
}

func readCsv(w http.ResponseWriter, r *http.Request) {
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("file") // Retrieve the file from form data

	if err != nil {
		// todo: improve error handling
	}

	defer file.Close()

	ReadCsv(file)

	fmt.Fprintf(w, "csv", handler.Filename)
	fmt.Println("Endpoint Hit: read csv!")

}

func ReadCsv(f multipart.File) ([][]string, error) {
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	fmt.Println(lines)
	return lines, nil
}

func HandleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/csv", readCsv).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
