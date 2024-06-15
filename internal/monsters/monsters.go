package monsters

import (
	"browser-mmo-backend/internal/data"
	"math/rand"
)

func GenerateMonster(mm data.MonsterModel, user data.User) (data.GeneratedMonster, error) {
	allMonsters, err := mm.GetAll()
	if err != nil {
		return data.GeneratedMonster{}, err
	}

	randomMonster := allMonsters[rand.Intn(len(allMonsters))]

	monster := data.GeneratedMonster{
		Name:         randomMonster.Name,
		ImageURL:     randomMonster.ImageURL,
		Level:        user.Level - 1,
		Constitution: float32(user.Constitution) * 0.7,
		Dexterity:    float32(user.Dexterity) * 0.7,
		Intelligence: float32(user.Intelligence) * 0.7,
		Strength:     float32(user.Strength) * 0.7,
	}

	return monster, nil
}
