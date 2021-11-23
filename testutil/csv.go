package testutil

import (
	"testing"

	"github.com/hamg26/academy-go-q42021/infrastructure/datastore"
)

type myCSV struct {
	FakeError error
	Records   [][]string
}

func (mycsv *myCSV) Close() {}

func (mycsv *myCSV) FindAll() (error, [][]string) {
	if mycsv.FakeError != nil {
		return mycsv.FakeError, nil
	}
	return nil, mycsv.Records
}

func NewCsvMock(t *testing.T, fakeError error, testCase string) datastore.MyCSV {
	t.Helper()
	testCases := map[string][][]string{
		"SUCCESS": {
			{"1", "name1", "type1"},
			{"2", "name2", "type2"},
		},
		"INVALID_ID": {
			{"asd", "name1", "type1"},
			{"2", "name2", "type2"},
		},
		"NIL": nil,
	}
	return &myCSV{FakeError: fakeError, Records: testCases[testCase]}
}
