package database

import (
	"GOBootcamp/app/models"
)

type MyCSV interface {
	ReadCsvFile() ([]models.JsonPost, error)
}
