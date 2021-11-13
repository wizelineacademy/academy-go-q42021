package services

import "github.com/AndresCravioto/academy-go-q42021/loader"

type StandServices interface {
	Initialize()
	SearchStandId(standId string) *[]*loader.StandData
}
