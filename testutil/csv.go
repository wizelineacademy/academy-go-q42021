package testutil

import (
	"testing"

	"github.com/hamg26/academy-go-q42021/infrastructure/datastore"
)

func NewCsvMock(t *testing.T, records [][]string) *datastore.MyCSV {
	t.Helper()
	return &datastore.MyCSV{Filepath: "file.csv", File: nil, Records: records}
}
