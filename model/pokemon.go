package model

type Pokemon struct {
	Id   int
	Name string
}

type PokeMonsters struct {
	Pokemon []*Pokemon
}

func NewPokemon(id int, name string)  *Pokemon{
	p := new(Pokemon)
	p.Id = id
	p.Name = name

	return p
}

func (p Pokemon) GetId() int {
	return p.Id
}

func (p Pokemon) GetName() string {
	return p.Name
}

func (p *PokeMonsters) SearchNameById(idToFind int) string {
	for _, poke := range p.Pokemon {
		if poke.Id == idToFind {
			return poke.Name
		}
	}

	return "none"
}
