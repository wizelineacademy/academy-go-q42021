package main

import (
	"encoding/json"
	"net/http"
)

type Car struct {
	Id    int    `json:"id"`
	Year  int    `json:"year"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Color string `json:"color"`
}

var (
	cars []Car
)

func init() {
	cars = []Car{Car{Id: 10, Year: 2020, Brand: "Nissan", Model: "Versa", Color: "Black"}}
}

func GetCars(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(cars)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "json enc error"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}
