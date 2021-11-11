package reader

import (
	"encoding/csv"
	"errors"
	"github.com/smmd/academy-go-q42021/model"
	"os"
	"strconv"
)

func GetPokeMonstersFromFile(filePath string) (*model.PokeMonsters, error) {
	file, err := csvToObject(filePath)
	pokeMonsters := make([]*model.Pokemon, 0)

	if err != nil {
		return nil, errors.New("Error reading file.")
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		return nil, errors.New("Error getting file records.")
	}

	for _, line := range lines {
		pokeId, _ := strconv.Atoi(line[0])

		pokeMonsters = append(pokeMonsters, &model.Pokemon{
			Id:	  pokeId,
			Name: line[1],
		})
	}

	return &model.PokeMonsters{
		Pokemon: pokeMonsters,
	}, nil
}

func csvToObject(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("File can not be read.")
	}

	return file, nil
}
