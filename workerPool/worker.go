package workerPool

import (
	"encoding/csv"
	"github.com/AndresCravioto/academy-go-q42021/models"
	"io"
	"os"
	"strconv"
	"sync"
)

const FileName = "repositories/files/champions.csv"

type WorkerCharacteristics struct {
	isOdd          bool
	limitPerWorker int
	maxItems       int
}

type WorkerRequest struct {
	Parity      string `json:"type" validate:"enum"`
	NumberOfItems  int `json:"items" validate:"required"`
	ItemsPerWorker int `json:"items_per_worker" validate:"required"`
}

type WorkerResponse struct {
	Value      interface{}
	Err        error
	JobRequest WorkerRequest
}

type WorkerHandler struct {}

func NewChampionWorker() WorkerHandler {
	return WorkerHandler{}
}

func (workerHandler WorkerHandler) ChampionsWorkerPool(request WorkerRequest) WorkerResponse {
	result := make([]*models.Champion, 0)
	channelCsvError := make(chan error, 1)
	channelJobs := make(chan []string, request.ItemsPerWorker)
	channelResult := make (chan *models.Champion)
	isOdd := request.Parity == "odd"
	itemsPerWorker := request.ItemsPerWorker
	numberOfItems := request.NumberOfItems

	characteristics := &WorkerCharacteristics{
		isOdd,
		itemsPerWorker,
		numberOfItems,
	}

	file, err := os.Open(FileName)
	if err != nil {
		return WorkerResponse{
			Value: nil,
			Err: err,
			JobRequest: request,
		}
	}

	defer file.Close()
	csvFileReader := csv.NewReader(file)

	var waitGroup sync.WaitGroup

	workerCount := numberOfItems / itemsPerWorker
	waitGroup.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer waitGroup.Done()
			worker(channelJobs, channelResult, characteristics)
		}()
	}

	go func() {
		for {
			rStr, err := csvFileReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				channelCsvError <- err
				break
			}

			channelJobs <- rStr
		}

		close(channelJobs)
		close(channelCsvError)
	}()

	//for _, e := range er

	go func() {
		waitGroup.Wait()
		close(channelResult)
	}()

	for r := range channelResult {
		result = append(result, r)
	}

	return WorkerResponse{
		Value: result,
		Err: nil,
		JobRequest: request,
	}
}

func worker(channelJobs <-chan []string, channelResult chan<- *models.Champion, conditions *WorkerCharacteristics) {
	countItems := 0

	for {
		job, ok := <-channelJobs

		if !ok {
			return
		}

		if countItems == conditions.limitPerWorker {
			return
		}

		if conditions.maxItems == 0 {
			return
		}

		championId, _ := strconv.Atoi(job[0])
		if conditions.isOdd && championId%2 != 0 {
			continue
		}

		if !conditions.isOdd && championId%2 == 0 {
			continue
		}

		channelResult <- parseChampion(job)

		conditions.maxItems--

		countItems++
	}
}

func parseChampion(data []string)  *models.Champion {
	return &models.Champion{
		ChampionID: data[0],
		Name: data[1],
		Title: data[2],
		Lore: data[3],
	}
}