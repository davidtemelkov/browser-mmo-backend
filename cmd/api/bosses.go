package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) createBossHandler(c *gin.Context) {
	var input struct {
		Position     int    `json:"position"`
		Name         string `json:"name"`
		ImageURL     string `json:"imageUrl"`
		Lvl          int    `json:"lvl"`
		Constitution int    `json:"constitution"`
		Dexterity    int    `json:"dexterity"`
		Intelligence int    `json:"intelligence"`
		Strength     int    `json:"strength"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	boss := &data.Boss{
		Position:     input.Position,
		Name:         input.Name,
		ImageURL:     input.ImageURL,
		Lvl:          input.Lvl,
		Constitution: input.Constitution,
		Dexterity:    input.Dexterity,
		Intelligence: input.Intelligence,
		Strength:     input.Strength,
	}

	err := app.models.Bosses.Insert(boss)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	c.IndentedJSON(http.StatusCreated, boss)
}

func (app *application) getBossByPositionHandler(c *gin.Context) {
	postition := c.Param("position")

	boss, err := app.models.Bosses.Get(postition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, boss)
}
