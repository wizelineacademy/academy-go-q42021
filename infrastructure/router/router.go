package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/hamg26/academy-go-q42021/config"
	"github.com/hamg26/academy-go-q42021/interface/controllers"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	config.ReadConfig()

	if config.C.Logging == true {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	e.GET("/pokemons", func(context echo.Context) error { return c.GetPokemons(context) })
	e.GET("/pokemons/:id", func(context echo.Context) error { return c.GetPokemon(context) })
	return e
}
