package controller

import (
	"net/http"

	"github.com/hamg26/academy-go-q42021/domain/model"
	"github.com/hamg26/academy-go-q42021/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(c Context) error
}

func NewPokemonController(ps interactor.PokemonInteractor) PokemonController {
	return &pokemonController{ps}
}

func (uc *pokemonController) GetPokemons(c Context) error {
	var p []*model.Pokemon

	err, p := uc.pokemonInteractor.Get()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
