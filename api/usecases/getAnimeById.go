package usecases

import (
	animeI "bootCampApi/api/interfaces"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func GetAnimeById(id string) animeI.AnimeStruct {
	url := "https://api.jikan.moe/v3/anime/" + id
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	animeId, _ := result["mal_id"].(float64)
	title := strings.Replace(result["title"].(string), ",", "", -1)
	synopsis := strings.Replace(result["synopsis"].(string), ",", "", -1)
	studio := strings.Replace(result["studios"].([]interface{})[0].(map[string]interface{})["name"].(string), ",", "", -1)
	animeData := animeI.AnimeStruct{
		AnimeId:  int(animeId),
		Title:    title,
		Synopsis: synopsis,
		Studio:   studio,
	}
	return animeData
}
