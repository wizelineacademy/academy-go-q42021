package reader

import (
	"encoding/csv"
	"errors"
	"github.com/smmd/academy-go-q42021/model"
	"os"
)

func GetPokeMonstersFromFile(filePath string) (*model.PokeMonsters, error) {
	file, err := csvToObject(filePath)

	if err != nil {
		return nil, errors.New("Error reading file.")
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	//TODO: foreach to convert Pokemons
}

func csvToObject(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("File can not be read.")
	}

	file.Close()

	return file, nil
}