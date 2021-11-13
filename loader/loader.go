package loader

import (
	"encoding/csv"
	"io"
	"log"
)

type StandData struct {
	StandID        string  `json:"standId"`
	Stand         string  `json:"stand"`
	Power       string  `json:"power"`
	Speed       string  `json:"speed"`
	Range       string  `json:"range"`
	Persistence       string  `json:"persistence"`
	Precision       string  `json:"precision"`
	Development       string  `json:"development"`

}

func LoadData(r io.Reader) *[]*StandData {
	reader := csv.NewReader(r)

	ret := make([]*StandData, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		if err != nil {
			log.Println(err)
		}
		stand := &StandData{
			StandID:        row[0],
			Stand:         row[1],
			Power:       row[2],
			Speed: 		 row[3],
			Range:          row[4],
			Persistence:        row[5],
			Precision:  		row[6],
			Development:      	row[7],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, stand)
	}
	return &ret
}
