package main

import (
	"bootCampApi/api"
	"log"
	"net/http"
	"os"
)

func main() {
	router := api.New()
	http.Handle("/", router)
	log.Println("Api running on port 8080")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
}
