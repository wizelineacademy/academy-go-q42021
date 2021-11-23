package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	models "github.com/hamg26/academy-go-q42021/domain/model"
)

const apiurl = "https://pokeapi.co/api/v2/"

func GetPokemon(id string) (err error, result models.PokemonDetails) {
	err = request(fmt.Sprintf("pokemon/%s", id), &result)
	return err, result
}

func request(endpoint string, obj interface{}) error {
	req, err := http.NewRequest(http.MethodGet, apiurl+endpoint, nil)
	if err != nil {
		return err
	}
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &obj)
}
