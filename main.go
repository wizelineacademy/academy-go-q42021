package main

import (
	"fmt"
	"log"
	router "mainRoot/infraestructure/router"
	controller "mainRoot/interface/controller"
	http "net/http"
)

func main() {
	controller.LlenarPokedex()
	router.Listen()

	fmt.Println("Servidor listo escuchando en " + router.Port)
	log.Fatal(http.ListenAndServe(router.Port, nil))
}
