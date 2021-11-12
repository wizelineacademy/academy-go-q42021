package entity

type Item struct {
	Id   int 	`json:"id"`
	Name string `json:"name"`
}

type ItemRepository interface {
	FetchItems() ([]Item, error)
	FetchItemByID(ID string) (*Item, error)
}