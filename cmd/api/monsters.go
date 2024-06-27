package main

import (
	"browser-mmo-backend/constants"
	"browser-mmo-backend/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) createMonsterHandler(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		ImageURL string `json:"imageURL"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	monster := &data.Monster{
		ID:       uuid.New().String(),
		Name:     input.Name,
		ImageURL: input.ImageURL,
	}

	err := app.models.Monsters.Insert(monster)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	c.IndentedJSON(http.StatusCreated, monster)
}
