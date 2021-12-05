package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Items - Description here....
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func writeCSV(str string) error {
	//str := `[{"id": 1,"name": "Pancho Lopez"},{"id": 2,"name": "Juana Juarez"},{"id": 4,"name": "Lola Contreras"},{"id": 6,"name": "Tomasa Perez"}]`
	var itemsList []Item

	if err := json.Unmarshal([]byte(str), &itemsList); err != nil {
		panic(err)
	}
	fmt.Println(itemsList)

	outputFile, err := os.Create("output.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)

	defer writer.Flush()

	for _, r := range itemsList {
		var csvRow []string
		csvRow = append(csvRow, fmt.Sprint(r.ID), r.Name)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	response, err := http.Get("http://localhost:8080/items")
	if err != nil {
		fmt.Printf("HTTP Fail with error %s\n", err)

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		writeCSV(string(data))
	}

}
