package model

type PokemonType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonTypeSlot struct {
	Slot int         `json:"slot"`
	Type PokemonType `json:"type"`
}

type PokemonAbility struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonAbilities struct {
	Ability  PokemonAbility `json:"ability"`
	IsHidden bool           `json:"is_hidden"`
	Slot     int            `json:"slot"`
}

type PokemonDetails struct {
	Height                 int                `json:"height"`
	Id                     uint64             `json:"id"`
	IsDefault              bool               `json:"is_default"`
	Name                   string             `json:"name"`
	Order                  int                `json:"order"`
	Weight                 int                `json:"weight"`
	BaseExperience         int                `json:"base_experience"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Types                  []PokemonTypeSlot  `json:"types"`
	Abilities              []PokemonAbilities `json:"abilities"`
}
