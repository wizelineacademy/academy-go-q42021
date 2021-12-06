package reader

import (
	"encoding/csv"
	"github.com/AndresCravioto/academy-go-q42021/models"
	"log"
	"os"
)

func GetAllStands(filePath string) (models.StandsInformation, error) {
	rows, err := csvToObject(filePath)

	stands := []models.Stand{}
	StandsInformation := models.StandsInformation{stands}
	if err != nil {
		log.Println(err)
		return StandsInformation, err
	}

	for _, row := range rows {
		stand := models.Stand{
			row[0],
			row[1],
			row[2],
			row[3],
			row[4],
			row[5],
			row[6],
			row[7],
		}

		StandsInformation.AddStand(stand)
	}

	return StandsInformation, nil
}


func SearchStandId(stands models.StandsInformation, standId string) []models.Stand {
	stand := Filter(stands.Stands, func(stand models.Stand) bool {
		return stand.StandID == standId
	})

	return stand
}

func Filter(stands []models.Stand, function func(models.Stand) bool) []models.Stand {

	filteredStands := []models.Stand{}
	for _, stand := range stands {
		if function(stand) {
			filteredStands = append(filteredStands, stand)
		}
	}
	return filteredStands
}

func csvToObject(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return lines, nil
}