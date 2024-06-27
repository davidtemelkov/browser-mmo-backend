package monsters

import (
	"browser-mmo-backend/data"
	"math"
	"math/rand"
)

// TODO: Balance mosters off of user lvl
func GenerateMonster(mm data.MonsterModel, user data.User) (data.GeneratedMonster, error) {
	allMonsters, err := mm.GetAll()
	if err != nil {
		return data.GeneratedMonster{}, err
	}

	randomMonster := allMonsters[rand.Intn(len(allMonsters))]

	monster := data.GeneratedMonster{
		Name:         randomMonster.Name,
		ImageURL:     randomMonster.ImageURL,
		Lvl:          user.Lvl - 1,
		Constitution: int(math.Round(float64(user.Constitution) * 0.6)),
		Dexterity:    int(math.Round(float64(user.Dexterity) * 0.6)),
		Intelligence: int(math.Round(float64(user.Intelligence) * 0.6)),
		Strength:     int(math.Round(float64(user.Strength) * 0.6)),
	}

	return monster, nil
}
