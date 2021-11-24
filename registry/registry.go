package registry

import (
	clients "github.com/hamg26/academy-go-q42021/infrastructure/clients"
	datatstore "github.com/hamg26/academy-go-q42021/infrastructure/datastore"
	controller "github.com/hamg26/academy-go-q42021/interface/controllers"
)

type registry struct {
	mycsv datatstore.MyCSV
	api   clients.ApiClient
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(mycsv datatstore.MyCSV, api clients.ApiClient) Registry {
	return &registry{mycsv, api}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
