package common

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"sync"

	"gobootcamp/models"
)

func CsvToPokemon(f multipart.File) (models.Pokemons, error) {
	var pokemons models.Pokemons

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return models.Pokemons{}, err
	}

	for _, item := range lines {
		pokemon := parsePokemon(item)

		pokemons = append(pokemons, pokemon)
	}

	fmt.Println(pokemons)
	return pokemons, nil
}

func worker(t string, jobs <-chan []string, results chan<- models.Pokemon) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}

			p := parsePokemon(job)
			if t == "odd" && p.Id%2 == 0 {
				results <- p
			} else if t == "even" && p.Id%2 != 0 {
				results <- p
			}
		}
	}
}

func WorkerPoolReadCSV(f multipart.File, numJobs int, itemsPerWorker int, t string) (models.Pokemons, error) {
	reader := csv.NewReader(f)
	var pokemons models.Pokemons

	numWorkers := numJobs / itemsPerWorker
	fmt.Println(numWorkers)
	jobs := make(chan []string, numJobs)
	res := make(chan models.Pokemon, numJobs)

	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(t, jobs, res)
		}()
	}

	for j := 1; j <= numJobs; j++ {
		rStr, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR: ", err.Error())
			break
		}
		jobs <- rStr
	}

	close(jobs)
	wg.Wait()
	close(res)

	for r := range res {
		pokemons = append(pokemons, r)
	}

	return pokemons, nil
}

func parsePokemon(data []string) models.Pokemon {
	id, _ := strconv.Atoi(data[0])
	pokemon := models.Pokemon{
		Id:   id,
		Name: data[1],
	}

	return pokemon
}
