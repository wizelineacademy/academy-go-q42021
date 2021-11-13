package app

import (
	"GOBootcamp/app/models"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to POST API")
	}
}

func (a *App) GetPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		f, err := os.Open("test.csv")
		if err != nil {
			log.Println("Unable to read input file test.csv", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Println("Unable to close file as CSV for ../../test.csv", err)
			}
		}(f)

		csvReader := csv.NewReader(f)
		records, err := csvReader.ReadAll()
		if err != nil {
			log.Println("Unable to parse file as CSV for ../../test.csv", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonPost, len(records))
		for row, content := range records {

			articleID, err := strconv.ParseUint(content[0], 10, 64)
			if err != nil {
				log.Println("Unable to parse file as CSV for ../../test.csv", err)
				sendResponse(w, r, nil, http.StatusInternalServerError)
				return
			}

			p := &models.Posts{
				ArticleID: articleID,
				Title:     content[1],
				Content:   content[2],
				Author:    content[3],
			}

			resp[row] = mapPostToJSON(p)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}
