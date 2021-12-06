package services

import (
	"encoding/json"
	standReader "github.com/AndresCravioto/academy-go-q42021/reader"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const FileName = "reader/fixtures/stands.csv"

func GetAllStands(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
		stands, err := standReader.GetAllStands(FileName)
	log.Println(stands, err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error fetching data"}`))
			return
		}

		b, err := json.Marshal(stands)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error marshalling data"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
}

func SearchByStandId(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if standId, ok := pathParams["standId"]; ok {
		log.Println(standId, ok)
		stands, err := standReader.GetAllStands(FileName)
		data := standReader.SearchStandId(stands, standId)

		b, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error marshalling data"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}