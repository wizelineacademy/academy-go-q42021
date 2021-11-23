package datastore

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/hamg26/academy-go-q42021/config"
)

type MyCSV interface {
	FindAll() (error, [][]string)
	Save([]string) error
}

type myCSV struct {
	Filepath string
	File     *os.File
	Records  [][]string
}

func closeFile(f *os.File) error {
	log.Println("Closing file")
	err := f.Close()
	if err != nil {
		log.Fatalln("Unable to close file", err)
		return err
	}
	return nil
}

func (mycsv *myCSV) FindAll() (error, [][]string) {
	if mycsv.Records != nil {
		log.Println("Returning cached records", mycsv.Filepath)
		return nil, mycsv.Records
	}

	log.Println("Reading records", mycsv.Filepath)
	f, err := os.Open(mycsv.Filepath)
	if err != nil {
		return err, nil
	}

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	mycsv.Records = records

	if err := closeFile(f); err != nil {
		return err, nil
	}

	return err, records
}

func (mycsv *myCSV) Save(record []string) error {
	log.Println("Saving record", record, mycsv.Filepath)

	f, err := os.Create(mycsv.Filepath)
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)

	records := append(mycsv.Records, record)

	if w.WriteAll(records) == nil {
		mycsv.Records = append(mycsv.Records, record)
		w.Flush()
	} else {
		log.Fatalln("Error saving record", err)
	}

	if err := closeFile(f); err != nil {
		return err
	}

	return err
}

func NewCSV() MyCSV {
	fp := config.C.CSV.Path
	return &myCSV{Filepath: fp}
}
