package main

import (
	"GOBootcamp/app"
	"log"
	"net/http"
	"os"
)

func main() {
	app := app.New()

	http.HandleFunc("/", app.Router.ServeHTTP)
	log.Println("App running...")

	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(0)
	}
}
