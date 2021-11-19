package router

import (
	controller "github.com/hamg26/academy-go-q42021/interface/controllers"
	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.GET("/pokemons", func(context echo.Context) error { return c.GetPokemons(context) })

	return e
}
