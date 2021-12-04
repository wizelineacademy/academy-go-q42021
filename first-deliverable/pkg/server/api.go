package server

import ( 
	"fmt"
	"encoding/json"
	"net/http"
	entity "academy-go-q42021/pkg/entity"
	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
	repository entity.ItemRepository
}

// Server - Entity to handle requests
type Server interface {
	Router() http.Handler
}

// New - Create a new entity
func New(repo entity.ItemRepository) Server {
	a:= &api{repository: repo}
	r:= mux.NewRouter()

	r.HandleFunc("/", a.homePage).Methods(http.MethodGet)
	r.HandleFunc("/items", a.fetchItems).Methods(http.MethodGet)
    r.HandleFunc("/item/{id}", a.fetchItem).Methods(http.MethodGet)

	a.router = r
	return a
}

// Router - Returns a router entity
func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func (a *api) fetchItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: fetchItems")
	records, err := a.repository.FetchItems()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound) 
		json.NewEncoder(w).Encode("Item Not found")
		return
	}
	json.NewEncoder(w).Encode(records)
}

func (a *api) fetchItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: fetchItem")
	vars := mux.Vars(r)
	item, err := a.repository.FetchItemByID(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Item Not found")
		return
	}
	json.NewEncoder(w).Encode(item)
}
