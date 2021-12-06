package repositories

import (
	"encoding/csv"
	"fmt"
	"github.com/AndresCravioto/academy-go-q42021/models"
	"log"
	"os"
)

type AllChampions struct {}

type ChampionsWriter struct {}

func CreateAllChampionsList() AllChampions {
	return AllChampions{}
}

func NewChampionsWriter() ChampionsWriter {
	return ChampionsWriter{}
}

func (AllChampions) GetAllChampions(filePath string) (models.ChampionsInformation, error) {
	rows, err := csvToStruct(filePath)

	champions := []models.Champion{}
	ChampionsInformation := models.ChampionsInformation{champions}

	if err != nil {
		log.Println(err)
		return ChampionsInformation, err
	}

	for _, row := range rows {
		champion := models.Champion{
			row[0],
			row[1],
			row[2],
			row[3],
		}

		ChampionsInformation.AddChampion(champion)
	}

	return ChampionsInformation, nil
}


func csvToStruct(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return lines, nil
}

func (ChampionsWriter) WriteChampionsInformation(response models.ChampionsApiResponse, filePath string) error {
	fmt.Println(response)
	file, err := os.OpenFile(filePath, os.O_RDWR, 0)

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, champion := range response.ChampionsData {
		champion := []string {champion.Key, champion.Name, champion.Title, champion.Blurb}

		err := writer.Write(champion)

		if err != nil {
			return err
		}
	}

	return nil
}