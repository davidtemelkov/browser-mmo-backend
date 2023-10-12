package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) createQuestandler(c *gin.Context) {
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
