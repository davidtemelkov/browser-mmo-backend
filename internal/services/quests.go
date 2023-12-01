package services

import (
	"browser-mmo-backend/internal/data"
	"math/rand"
)

// TODO: Refactor this to require user to make tailored quests
func GenerateQuestsForUser(allQuests []data.Quest) ([]data.GeneratedQuest, error) {
	var generatedQuests []data.GeneratedQuest
	for i := 0; i < 3; i++ {
		randIndex := rand.Intn(len(allQuests))
		selectedQuest := allQuests[randIndex]

		generatedQuest := data.GeneratedQuest{
			Name:     selectedQuest.Name,
			ImageURL: selectedQuest.ImageURL,
			EXP:      "10", //Add logic for calculating exp rewards
			Gold:     "10", //Add logic for calculating gold rewards
		}

		timeOptions := []string{"5 mins", "10 mins", "15 mins"}
		randTimeIndex := rand.Intn(len(timeOptions))
		generatedQuest.Time = timeOptions[randTimeIndex]

		generatedQuests = append(generatedQuests, generatedQuest)
		allQuests = append(allQuests[:randIndex], allQuests[randIndex+1:]...)
	}

	return generatedQuests, nil
}
