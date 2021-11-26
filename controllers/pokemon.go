package controllers

import (
	"encoding/json"
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
		//common.HandleInternalServerError(w)
	}

	p.PokemonRepo.SaveManyPokemons(pokemons)
	jsonResp, err := json.Marshal(pokemons)

	if err != nil {
		//common.HandleInternalServerError(w)
	}

	c.JSON(http.StatusCreated, jsonResp)
}

func (p *PokemonController) GetPokemonById(c *gin.Context) {
	// question: is there a simpliest way to parse the param?
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	pokemon, err := p.PokemonRepo.GetPokemonById(int(id))

	if err != nil {
		//common.HandleInternalServerError(w)
	}

	jsonResp, err := json.Marshal(pokemon)

	if err != nil {
		//common.HandleInternalServerError(w)
	}

	c.JSON(http.StatusOK, jsonResp)
}
