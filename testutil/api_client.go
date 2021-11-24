package testutil

import (
	"github.com/hamg26/academy-go-q42021/domain/model"

	"github.com/stretchr/testify/mock"
)

type ApiClient struct {
	mock.Mock
}

func (ac ApiClient) GetPokemon(id string) (error, *model.PokemonDetails) {
	args := ac.Called(id)
	if args.Get(0) != nil {
		return args.Error(1), args.Get(0).(*model.PokemonDetails)
	}
	return args.Error(1), nil
}
