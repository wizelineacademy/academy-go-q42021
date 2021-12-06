package services

import (
	"github.com/AndresCravioto/academy-go-q42021/models"
	"github.com/gorilla/mux"
	"net/http"
)

const FileName = "repositories/files/champions.csv"

type getter interface {
	GetAllChampions(filePath string) (models.ChampionsInformation, error)
}

type SearchService struct {
	repository getter
}

func NewSearchService(repository getter) SearchService {
	return SearchService{repository}
}

func (searchService SearchService) GetAllChampions() (models.ChampionsInformation, error) {
	return searchService.repository.GetAllChampions(FileName)
}

func (searchService SearchService) SearchChampionById(request *http.Request) ([]models.Champion, error) {
	pathParams := mux.Vars(request)
	championId := pathParams["championId"]
	championsInformation, err := searchService.repository.GetAllChampions(FileName)

	if err != nil {
		return *new([]models.Champion), err
	}

	champions := Filter(championsInformation.Champions, func(champion models.Champion) bool {
		return champion.ChampionID == championId
	})

	return champions, nil
}

func Filter(champions []models.Champion, function func(models.Champion) bool) []models.Champion {

	filteredChampions := []models.Champion{}
	for _, champion := range champions {
		if function(champion) {
			filteredChampions = append(filteredChampions, champion)
		}
	}
	return filteredChampions
}