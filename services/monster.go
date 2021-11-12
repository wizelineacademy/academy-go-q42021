package services

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/Le-MaliX/ACADEMY-GO-Q42021/models"
)

func openCsv() ([][]string, error) {
	file, err := os.Open("./data/monsters.csv")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	f := csv.NewReader(file)
	f.Comma = ';'
	lines, err := f.ReadAll()
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func GetAllMonsters() ([]models.Monster, error) {
	csvLines, err := openCsv()
	if err != nil {
		return nil, err
	}

	var monsters []models.Monster

	for _, line := range csvLines {
		monster := models.Monster{
			Id:              line[0],
			Name:            line[1],
			ChallengeRating: line[2],
			HPDice:          line[18],
			HP:              line[19],
		}
		monsters = append(monsters, monster)
	}

	return monsters[1:], nil
}

func GetMonsterById(id string) (models.Monster, error) {
	var monster models.Monster
	csvLines, err := openCsv()
	if err != nil {
		return monster, err
	}

	for _, line := range csvLines {
		if line[0] != id {
			continue
		}
		monster = models.Monster{
			Id:              line[0],
			Name:            line[1],
			ChallengeRating: line[2],
			HPDice:          line[18],
			HP:              line[19],
		}
		return monster, nil
	}

	err = fmt.Errorf("couldn't find id: %v", id)

	return monster, err
}
