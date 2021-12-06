package main

import (
standRouter "github.com/AndresCravioto/academy-go-q42021/api/routes"
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
}

func main() {
	standRouter.Routes()
}