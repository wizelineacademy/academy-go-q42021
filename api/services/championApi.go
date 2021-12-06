package services

import (
	"encoding/json"
	"fmt"
	"github.com/AndresCravioto/academy-go-q42021/models"
	"io/ioutil"
	"net/http"
)

type setter interface {
	WriteChampionsInformation(response models.ChampionsApiResponse, filePath string) error
}

type WriteService struct {
	repository setter
}

func NewWriteService(repository setter) WriteService {
	return WriteService{repository}
}

func (writeService WriteService) ConsumeDDragonChampionsApi() error {
	response, err := http.Get("https://ddragon.leagueoflegends.com/cdn/11.23.1/data/en_US/champion.json")

	if err != nil {
		return err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseObject models.ChampionsApiResponse

	json.Unmarshal(responseData, &responseObject)
	fmt.Println(err)
	err = writeService.repository.WriteChampionsInformation(responseObject, FileName)

	if err != nil {
		return err
	}

	return nil
}