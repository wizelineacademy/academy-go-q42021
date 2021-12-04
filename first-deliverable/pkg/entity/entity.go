package entity

// Items - Description here....
type Item struct {
	ID   int 	`json:"id"`
	Name string `json:"name"`
}

// ItemRepository - Description here....
type ItemRepository interface {
	FetchItems() ([]Item, error)
	FetchItemByID(ID string) (*Item, error)
}