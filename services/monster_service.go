package services

import (
	"fmt"

	"github.com/Le-MaliX/ACADEMY-GO-Q42021/domain/model"
	"github.com/Le-MaliX/ACADEMY-GO-Q42021/interface/repository"
)

func GetAllMonsters() ([]model.Monster, error) {
	monsters, err := repository.GetMonstersData()
	if err != nil {
		return nil, err
	}

	return monsters, nil
}

func GetMonsterById(id string) (model.Monster, error) {
	var monster model.Monster
	monsters, err := repository.GetMonstersData()
	if err != nil {
		return monster, err
	}

	for _, monster := range monsters {
		if monster.Id != id {
			continue
		}
		return monster, nil
	}

	err = fmt.Errorf("couldn't find id: %v", id)

	return monster, err
}
