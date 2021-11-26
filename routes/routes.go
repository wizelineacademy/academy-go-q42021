package routes

import (
	"gobootcamp/controllers"
	"gobootcamp/repositories"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/_health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "app is working!",
		})
	})

	pokemonController := controllers.PokemonController{PokemonRepo: &repositories.PokemonRepository{}}
	r.GET("/pokemon/:id", pokemonController.GetPokemonById)
	r.POST("/pokemon/csv", pokemonController.ReadCsv)

	r.Run() // listen and serve on 0.0.0.0:8080
}
