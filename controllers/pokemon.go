package controllers

import (
	"net/http"
	"strconv"

	"gobootcamp/common"
	"gobootcamp/repositories"

	"github.com/gin-gonic/gin"
)

type PokemonController struct {
	PokemonRepo *repositories.PokemonRepository
}

func (p *PokemonController) ReadCsv(c *gin.Context) {

	fileHeader, _ := c.FormFile("file")
	file, _ := fileHeader.Open()

	pokemons, err := common.CsvToPokemon(file)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "csv not well formated"})
	}

	p.PokemonRepo.SaveManyPokemons(pokemons)

	c.JSON(http.StatusCreated, pokemons)
}

func (p *PokemonController) GetPokemonById(c *gin.Context) {
	// question: is there a simpliest way to parse the param?
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	pokemon, err := p.PokemonRepo.GetPokemonById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pokemon not found"})
	}

	c.JSON(http.StatusOK, pokemon)
}
