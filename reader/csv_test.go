package reader

import (
	"github.com/smmd/academy-go-q42021/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertingCSVDataToModelObj(t *testing.T)  {
	expected := &model.PokeMonsters{
		[]*model.Pokemon{
			{
				Id: 1,
				Name: "bulbasaur",
			},
			{
				Id: 2,
				Name: "ivysaur",
			},
			{
				Id: 3,
				Name: "venusaur",
			},
		},
	}

	actual, _ := GetPokeMonstersFromFile("fixtures/pokedex_data.csv")

	assert.Equal(t, actual, expected)
}

func TestThrowingErrorFileNoExist(t *testing.T)  {
	_, actual := GetPokeMonstersFromFile("fixtures/pokedex_data_fail.csv")

	assert.EqualError(t, actual, "Error reading file.")
}
