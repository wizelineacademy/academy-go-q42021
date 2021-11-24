package main

import (
	"io/ioutil"
	"log"

	"github.com/labstack/echo"

	"github.com/hamg26/academy-go-q42021/config"
	clients "github.com/hamg26/academy-go-q42021/infrastructure/clients"
	"github.com/hamg26/academy-go-q42021/infrastructure/datastore"
	"github.com/hamg26/academy-go-q42021/infrastructure/router"
	"github.com/hamg26/academy-go-q42021/registry"
)

func main() {
	config.ReadConfig()

	if config.C.Logging != true {
		log.SetOutput(ioutil.Discard)
		log.SetFlags(0)
	}

	mycsv := datastore.NewCSV()
	api := clients.NewApiClient("https://pokeapi.co/api/v2/")

	r := registry.NewRegistry(mycsv, api)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatal(err)
	}
}
