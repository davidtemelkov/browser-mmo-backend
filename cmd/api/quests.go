package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"errors"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (app *application) createQuestHandler(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		ImageURL string `json:"imageURL"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError.Error())
		return
	}

	quest := &data.Quest{
		ID:       uuid.New().String(),
		Name:     input.Name,
		ImageURL: input.ImageURL,
	}

	err := app.models.Quests.Insert(quest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, quest)
}

func (app *application) generateQuestsHandler(c *gin.Context) {
	email := c.Param("email")

	claims, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	userEmail, exists := claims.(jwt.MapClaims)["email"].(string)
	if !exists {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	_, err := app.models.Users.Get(email)
	if err != nil {
		if errors.Is(err, constants.UserNotFoundError) {
			c.JSON(http.StatusNotFound, constants.UserNotFoundError.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	if email != userEmail {
		c.JSON(http.StatusForbidden, constants.UserIsNotAuthorizedError.Error())
		return
	}

	allQuests, err := app.models.Quests.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

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

	err = app.models.Users.AddGeneratedQuests(email, generatedQuests)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, generatedQuests)
}
