package testutil

import (
	"testing"

	"github.com/hamg26/academy-go-q42021/domain/model"
	pi "github.com/hamg26/academy-go-q42021/usecase/interactor"
)

type pokemonInteractor struct {
	FakeError error
	Records   []*model.Pokemon
}

func (pi *pokemonInteractor) GetOne(id uint64) (error, *model.Pokemon) {
	if pi.Records == nil {
		return pi.FakeError, nil
	}
	return pi.FakeError, pi.Records[0]
}

func (pi *pokemonInteractor) GetAll() (error, []*model.Pokemon) {
	return pi.FakeError, pi.Records
}

func NewPokemonInteractorMock(t *testing.T, fakeError error, testCase string) pi.PokemonInteractor {
	t.Helper()
	testCases := map[string][]*model.Pokemon{
		"SUCCESS": {
			&model.Pokemon{Id: uint64(1), Name: "name1", Type: "type1"},
			&model.Pokemon{Id: uint64(2), Name: "name2", Type: "type2"},
		},
		"NIL": nil,
	}
	return &pokemonInteractor{FakeError: fakeError, Records: testCases[testCase]}
}
