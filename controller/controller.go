package controller

import (
	"encoding/json"
	"fmt"
	"github.com/AndresCravioto/academy-go-q42021/models"
	"github.com/AndresCravioto/academy-go-q42021/workerPool"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type search interface {
	GetAllChampions() (models.ChampionsInformation, error)
	SearchChampionById(request *http.Request) ([]models.Champion, error)
}

type ddragonChampionApi interface {
	ConsumeDDragonChampionsApi() error
}

type championsWorker interface {
	ChampionsWorkerPool(workerPool.WorkerRequest) workerPool.WorkerResponse
}

type ChampionsHandler struct {
	searchService search
	apiService ddragonChampionApi
	championsWorker championsWorker
}

func NewChampionsHandler(search search, ddragonChampionApi ddragonChampionApi, championsWorker championsWorker) ChampionsHandler {
	return ChampionsHandler {
		search,
		ddragonChampionApi,
		championsWorker,
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

func (championsHandler ChampionsHandler) DDragonChampionsList(writer http.ResponseWriter, request *http.Request) {
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

func (championsHandler ChampionsHandler) ChampionsByWorker(writer http.ResponseWriter, request *http.Request)  {
	pathParams := mux.Vars(request)
	parity := pathParams["type"]
	numberOfItems := pathParams["items"]
	itemsPerWorkers := pathParams["items_per_worker"]

	workerRequest, err := workerRequest(
		parity,
		numberOfItems,
		itemsPerWorkers,
	)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	response := championsHandler.championsWorker.ChampionsWorkerPool(workerRequest)
	marshalledResponse, err := json.Marshal(response)

	if response.Err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(response.Err.Error()))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshalledResponse)
}

func workerRequest(parity string, numberOfItems string, itemsPerWorker string) (workerPool.WorkerRequest, error){
	numOfItems, _ := strconv.Atoi(numberOfItems)
	numberOfItemsPerWorker, _ := strconv.Atoi(itemsPerWorker)
	v := validator.New()
	_ = v.RegisterValidation("enum", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "odd" || fl.Field().String() == "even"
	})

	request := workerPool.WorkerRequest{
		Parity: parity,
		NumberOfItems: numOfItems,
		ItemsPerWorker: numberOfItemsPerWorker,
	}

	err := v.Struct(request)
	if err != nil {
		return request, fmt.Errorf("invalid request: %w", err)
	}

	return request, nil
}