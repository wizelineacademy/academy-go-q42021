package entity

// Items - Description here....
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ItemRepository - Description here....
type ItemRepository interface {
	FetchItems() ([]Item, error)
	FetchItemByID(ID string) (*Item, error)
}

type JSONOutput struct {
	Name           string         `json:"name"`
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
}

type Description struct {
	Description string `json:"description"`
	Language    Region `json:"language"`
}

type Region struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Name struct {
	Language Region `json:"language"`
	Name     string `json:"name"`
}

type PokemonEntry struct {
	EntryNumber    int64  `json:"entry_number"`
	PokemonSpecies Region `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}
