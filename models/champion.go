package models

type Champion struct {
	ChampionID		string  `json:"championId"`
	Name       string  `json:"champion"`
	Title       string  `json:"title"`
	Lore       string  `json:"lore"`
}

type ChampionsInformation struct {
	Champions []Champion
}


func NewChampion(id string, name string, title string, lore string) *Champion{
	champion := new(Champion)
	champion.ChampionID = id
	champion.Name = name
	champion.Title = title
	champion.Lore = lore

	return champion
}

func (championsList *ChampionsInformation) AddChampion(champion Champion) []Champion {
	championsList.Champions = append(championsList.Champions, champion)

	return championsList.Champions
}

func (champion Champion) GetId() string {
	return champion.ChampionID
}

func (champion *Champion) SetId(id string) {
	champion.ChampionID = id
}

func (champion Champion) GetName() string {
	return champion.Name
}

func (champion *Champion) SetName(name string) {
	champion.Name = name
}

func (champion Champion) GetTitle() string {
	return champion.Title
}

func (champion *Champion) SetTitle(title string) {
	champion.Title = title
}

func (champion Champion) GetLore() string {
	return champion.Lore
}

func (champion *Champion) SetLore(lore string) {
	champion.Lore = lore
}