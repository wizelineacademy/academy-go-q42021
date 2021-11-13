package repository

import (
	"encoding/csv"
	"os"

	"github.com/Le-MaliX/ACADEMY-GO-Q42021/domain/model"
)

func openCsv() ([][]string, error) {
	file, err := os.Open("./infrastructure/datastore/monsters.csv")
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

func GetMonstersData() ([]model.Monster, error) {
	csvLines, err := openCsv()
	if err != nil {
		return nil, err
	}

	var monsters []model.Monster

	for _, line := range csvLines {
		monster := model.Monster{
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
