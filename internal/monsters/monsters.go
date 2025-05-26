package monsters

import (
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/gamecontent"
	"math"
	"math/rand"
)

type GeneratedMonster struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Lvl          int    `json:"lvl"`
	Constitution int    `json:"constitution"`
	Dexterity    int    `json:"dexterity"`
	Intelligence int    `json:"intelligence"`
	Strength     int    `json:"strength"`
}

// TODO: Balance mosters off of user lvl, can also have a balance.json
func GenerateMonster(allMonsters []gamecontent.Monster, user data.User) (GeneratedMonster, error) {

	randomMonster := allMonsters[rand.Intn(len(allMonsters))]

	monster := GeneratedMonster{
		ID:           randomMonster.ID,
		Name:         randomMonster.Name,
		Lvl:          user.Lvl - 1,
		Constitution: int(math.Round(float64(user.Constitution) * 0.6)),
		Dexterity:    int(math.Round(float64(user.Dexterity) * 0.6)),
		Intelligence: int(math.Round(float64(user.Intelligence) * 0.6)),
		Strength:     int(math.Round(float64(user.Strength) * 0.6)),
	}

	return monster, nil
}
