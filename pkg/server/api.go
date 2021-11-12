package server


import ( 
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	entity "academy-go-q42021/pkg/entity"
)

type api struct {
	router http.Handler
	repository entity.ItemRepository
}

type Server interface {
	Router() http.Handler
}

func New(repo entity.ItemRepository) Server {
	a:= &api{repository: repo}
	r:= mux.NewRouter()

	r.HandleFunc("/", a.homePage).Methods(http.MethodGet)
	r.HandleFunc("/items", a.fetchItems).Methods(http.MethodGet)
    r.HandleFunc("/item/{id}", a.fetchItem).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func (a *api) fetchItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: fetchItems")
	records, _ := a.repository.FetchItems()

	w.Header().Set("Content-Type", "application/json")
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
