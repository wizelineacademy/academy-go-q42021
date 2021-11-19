package interactor

import (
	"github.com/hamg26/academy-go-q42021/domain/model"
	"github.com/hamg26/academy-go-q42021/usecase/presenter"
	"github.com/hamg26/academy-go-q42021/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	Get() (error, []*model.Pokemon)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (ps *pokemonInteractor) Get() (error, []*model.Pokemon) {
	err, p := ps.PokemonRepository.FindAll()

	return err, ps.PokemonPresenter.ResponsePokemons(p)
}
