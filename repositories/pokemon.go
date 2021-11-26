package repositories

import (
	"errors"

	"gobootcamp/models"
)

/* type PokemonRepository interface {
	// Question: for what is used/required the context? hos to pass it
	GetPokemonById(ctx context.Context, id int) (*models.Pokemon, error)
	SaveManyPokemons(ctx context.Context, pokemons []models.Pokemon)
} */

type PokemonRepository struct {
	pokemons []models.Pokemon
}

func (p *PokemonRepository) SaveManyPokemons(pokemons []models.Pokemon) {
	p.pokemons = pokemons
}

func (p *PokemonRepository) GetPokemonById(id int) (models.Pokemon, error) {
	for _, val := range p.pokemons {
		if val.Id == id {
			return val, nil
		}

	}

	// Question: is this right? passing an empty model
	return models.Pokemon{}, errors.New("pokemon not found")

}
