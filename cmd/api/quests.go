package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
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
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	allQuests, err := app.models.Quests.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	generatedQuests, err := services.GenerateQuestsForUser(allQuests)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	err = app.models.Quests.SetQuests(user.Email, generatedQuests)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, generatedQuests)
}

func (app *application) setCurrentQuestHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	var currentQuestMap map[string]data.GeneratedQuest
	if err := c.BindJSON(&currentQuestMap); err != nil {
		c.JSON(http.StatusBadRequest, constants.InvalidJSONFormatError.Error())
		return
	}

	//TODO: This could be in a single request
	err := app.models.Quests.SetCurrentQuest(user.Email, currentQuestMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	err = app.models.Quests.SetQuests(user.Email, []data.GeneratedQuest{
		{Name: "Empty Quest 0", ImageURL: "", Time: "", EXP: "0", Gold: "0"},
		{Name: "Empty Quest 1", ImageURL: "", Time: "", EXP: "0", Gold: "0"},
		{Name: "Empty Quest 2", ImageURL: "", Time: "", EXP: "0", Gold: "0"},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}
	//end

	c.IndentedJSON(http.StatusOK, currentQuestMap)
}

func (app *application) cancelCurrentQuestHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	err := app.models.Quests.CancelCurrentQuest(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "quest cancelled")
}
