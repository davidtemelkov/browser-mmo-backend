package quests

import (
	"browser-mmo-backend/internal/gamecontent"
	"math/rand"
)

type GeneratedQuest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Time string `json:"time"`
	EXP  string `json:"exp"`
	Gold string `json:"gold"`
}

// TODO: Refactor this to require user to make tailored quests
func GenerateQuestsForUser(allQuests []gamecontent.Quest) ([]GeneratedQuest, error) {
	var generatedQuests []GeneratedQuest
	for i := 0; i < 3; i++ {
		randIndex := rand.Intn(len(allQuests))
		selectedQuest := allQuests[randIndex]

		generatedQuest := GeneratedQuest{
			Name: selectedQuest.Name,
			ID:   selectedQuest.ID,
			EXP:  "10", //Add logic for calculating exp rewards
			Gold: "10", //Add logic for calculating gold rewards
		}

		timeOptions := []string{"5 mins", "10 mins", "15 mins"}
		randTimeIndex := rand.Intn(len(timeOptions))
		generatedQuest.Time = timeOptions[randTimeIndex]

		generatedQuests = append(generatedQuests, generatedQuest)
		allQuests = append(allQuests[:randIndex], allQuests[randIndex+1:]...)
	}

	return generatedQuests, nil
}
