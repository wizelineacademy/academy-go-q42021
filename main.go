package main

import (
	"fmt"
	pokereader "github.com/smmd/academy-go-q42021/reader"
)

func main() {
	var pokeId int
	var filePath string

	fmt.Print("Enter a Pokemon ID: ")

	_, err := fmt.Scanf("%d", &pokeId)

	if err != nil {
		panic(fmt.Errorf("could not read ID: %w", err))
	}

	fmt.Print("Enter the csv path: ")

	_, err = fmt.Scanf("%s", &filePath)

	if err != nil {
		panic(fmt.Errorf("could not read Path: %w", err))
	}

	pokeMonsters, err := pokereader.GetPokeMonstersFromFile(filePath)

	if err != nil {
		panic(err)
	}

	fmt.Println(pokeMonsters.SearchNameById(pokeId))
}
