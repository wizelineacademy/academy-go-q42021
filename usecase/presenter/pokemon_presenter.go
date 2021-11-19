package presenter

import "github.com/hamg26/academy-go-q42021/domain/model"

type PokemonPresenter interface {
	ResponsePokemons(u []*model.Pokemon) []*model.Pokemon
}
