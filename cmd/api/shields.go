package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) createShieldHandler(c *gin.Context) {
	var input struct {
		BaseName       string `json:"base_name"`
		MinLevel       int    `json:"min_level"`
		BlockChanceMin int    `json:"block_chance_min"`
		BlockChanceMax int    `json:"block_chance_max"`
		IsLegendary    bool   `json:"is_legendary"`
		ImageURL       string `json:"imageURL"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError.Error())
		return
	}

	shield := &data.Shield{
		ID:             uuid.New().String(),
		BaseName:       input.BaseName,
		MinLevel:       input.MinLevel,
		BlockChanceMin: input.BlockChanceMin,
		BlockChanceMax: input.BlockChanceMax,
		IsLegendary:    input.IsLegendary,
		ImageURL:       input.ImageURL,
	}

	err := app.models.Shields.Insert(shield)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, shield)
}
