package repository

import (
	"sync"
	"strconv"
	"fmt"		
	record "academy-go-q42021/pkg/entity"
)

type itemRepository struct {
	mtx     sync.RWMutex
	records []record.Item
}

// NewItemsRepository - Create repository of Items
func NewItemsRepository(records []record.Item) record.ItemRepository {
	if records == nil {
		records = make([]record.Item,0)
	}
	return &itemRepository{
		records: records,
	}
}

// FetchItems - Return all Items.
func (r *itemRepository) FetchItems() ([]record.Item, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	values := make([]record.Item, 0, len(r.records))
	for _, value := range r.records {
		values = append(values, value)
	}
	return values, nil
}

// FetchItemByID - Returns a specific item based in its id
func (r *itemRepository) FetchItemByID(ID string) (*record.Item, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	id , _ := strconv.Atoi(ID)
	fmt.Println("value of id" , id)
	for _, v := range r.records {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("The ID %s doesn't exist", ID)
}
