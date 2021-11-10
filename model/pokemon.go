package model

type Pokemon struct {
	id int
	name string
}

type PokeMonsters struct {
	Pokemon []Pokemon
}

func NewPokemon(id int, name string)  *Pokemon{
	p := new(Pokemon)
	p.id = id
	p.name = name

	return p
}

func (p Pokemon) GetId() int {
	return p.id
}

func (p Pokemon) GetName() string {
	return p.name
}
