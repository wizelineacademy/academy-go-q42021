package controller

import (
	"encoding/json"
	"fmt"
	"github.com/AndresCravioto/academy-go-q42021/models"
	"net/http"
)

type search interface {
	GetAllChampions() (models.ChampionsInformation, error)
	SearchChampionById(request *http.Request) ([]models.Champion, error)
}

type ddragonChampionApi interface {
	ConsumeDDragonChampionsApi() error
}

type ChampionsHandler struct {
	searchService search
	apiService ddragonChampionApi
}

func NewChampionsHandler(search search, ddragonChampionApi ddragonChampionApi) ChampionsHandler {
	return ChampionsHandler {
		search,
		ddragonChampionApi,
	}
}

func (championsHandler ChampionsHandler) ChampionsInformation(writer http.ResponseWriter, request *http.Request) {
	championsInformation, err := championsHandler.searchService.GetAllChampions()
	marshalledChampionsInformation, err := json.Marshal(championsInformation)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error": "error fetching data"}`))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshalledChampionsInformation)
}

func (championsHandler ChampionsHandler) Champion(writer http.ResponseWriter, request *http.Request) {
	champion, err := championsHandler.searchService.SearchChampionById(request)
	marshalledChampionInformation, err := json.Marshal(champion)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error": "error fetching data"}`))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshalledChampionInformation)
}

func (championsHandler ChampionsHandler) ChampionList(writer http.ResponseWriter, request *http.Request) {
	err := championsHandler.apiService.ConsumeDDragonChampionsApi()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error": "error fetching data"}`))
		return
	}

	response := make(map[string]string)
	fmt.Println(response)
	response["message"] = "OK"
	marshalledResponse, err := json.Marshal(response)


	writer.WriteHeader(http.StatusOK)
	writer.Write(marshalledResponse)
}