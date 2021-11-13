package app

import (
	"GOBootcamp/app/models"
	"encoding/json"
	"log"
	"net/http"
)

func mapPostToJSON(p *models.Posts) models.JsonPost {
	return models.JsonPost{
		ArticleID: p.ArticleID,
		Title:     p.Title,
		Content:   p.Content,
		Author:    p.Author,
	}
}

func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json response. err=%v\n", err)
	}
}
