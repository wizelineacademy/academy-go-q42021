package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewPokemon(t *testing.T) {
	poke := NewPokemon(3, "pikachu")
	expected := &Pokemon{
		3,
		"pikachu",
	}

	actual := poke
	assert.Equal(t, actual, expected)
}

func TestGettingPokemonNameById(t *testing.T)  {
	pokemonsers := &PokeMonsters{
		[]*Pokemon{
			{
				Id: 10,
				Name: "jigglypuff",
			},
			{
				Id: 150,
				Name: "mewtwo",
			},
		},
	}

	expected := "jigglypuff"
	actual := pokemonsers.SearchNameById(10)

	assert.Equal(t, actual, expected)
}
