package services

import (
	"github.com/AndresCravioto/academy-go-q42021/loader"
	"log"
	"os"
)

type Stands struct {
	Store *[]*loader.StandData `json:"store"`
}

func (b *Stands) Initialize() {
	filename := "./assets/stands.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	b.Store = loader.LoadData(file)
}

func (b *Stands) SearchStandId(standId string) *[]*loader.StandData {
	ret := Filter(b.Store, func(v *loader.StandData) bool {
		return v.StandID == standId
	})

	return ret
}

func Filter(vs *[]*loader.StandData, f func(*loader.StandData) bool) *[]*loader.StandData {
	vsf := make([]*loader.StandData, 0)
	for _, v := range *vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return &vsf
}
