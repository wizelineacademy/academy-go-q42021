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
	GetAll() (error, []*model.Pokemon)
	GetOne(id uint64) (error, *model.Pokemon)
	GetOneDetails(id string) (error, *model.PokemonDetails)
	SavePokemon(*model.PokemonDetails) error
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (ps *pokemonInteractor) GetAll() (error, []*model.Pokemon) {
	err, p := ps.PokemonRepository.FindAll()

	return err, ps.PokemonPresenter.ResponsePokemons(p)
}

func (ps *pokemonInteractor) GetOne(id uint64) (error, *model.Pokemon) {
	err, p := ps.PokemonRepository.FindOne(id)

	return err, ps.PokemonPresenter.ResponsePokemon(p)
}

func (ps *pokemonInteractor) GetOneDetails(id string) (error, *model.PokemonDetails) {
	err, p := ps.PokemonRepository.FindOneDetails(id)

	return err, ps.PokemonPresenter.ResponsePokemonDetails(p)
}

func (ps *pokemonInteractor) SavePokemon(p *model.PokemonDetails) error {
	err := ps.PokemonRepository.SavePokemon(p)

	return err
}
