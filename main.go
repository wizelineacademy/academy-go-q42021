package main

import (
	"fmt"
	pokereader "github.com/smmd/academy-go-q42021/reader"
)

func main() {
	var pokeId int

	fmt.Print("Enter a Pokemon ID: ")
	fmt.Scanf("%d", &pokeId)

	//TODO: second iteration input file path
	pokemonsters, err := pokereader.GetPokeMonstersFromFile("pokedex_data.csv")

	if err != nil {
		panic(err)
	}

	fmt.Println(pokemonsters.SearchNameById(pokeId))
}
