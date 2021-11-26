package controller

type pokemonController interface {
	GetPokemons(c Context) error
	GetPokemon(c Context) error
	GetPokemonDetails(c Context) error
}

type AppController interface {
	pokemonController
}
