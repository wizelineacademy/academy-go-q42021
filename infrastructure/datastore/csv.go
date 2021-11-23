package datastore

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/hamg26/academy-go-q42021/config"
)

type MyCSV interface {
	FindAll() (error, [][]string)
	Close()
}

type myCSV struct {
	Filepath string
	File     *os.File
	Records  [][]string
}

func readfile(filepath string) (error, *os.File) {
	log.Println("Reading", filepath)
	f, err := os.Open(filepath)
	return err, f
}

func (mycsv *myCSV) Close() {
	log.Println("Closing", mycsv.Filepath)
	err := mycsv.File.Close()
	if err != nil {
		log.Fatalln("Unable to close csv", err)
	}
}

func (mycsv *myCSV) FindAll() (error, [][]string) {
	if mycsv.Records != nil {
		log.Println("Returning cached records", mycsv.Filepath)
		return nil, mycsv.Records
	}

	log.Println("Reading records", mycsv.Filepath)
	csvReader := csv.NewReader(mycsv.File)
	records, err := csvReader.ReadAll()
	mycsv.Records = records
	return err, records
}

func NewCSV() (error, MyCSV) {
	fp := config.C.CSV.Path
	err, f := readfile(fp)
	return err, &myCSV{Filepath: fp, File: f}
}
