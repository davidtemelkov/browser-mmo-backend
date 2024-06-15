package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) createArmourHandler(c *gin.Context) {
	var input struct {
		BaseName        string `json:"base_name"`
		Type            string `json:"type"`
		MinLevel        int    `json:"min_level"`
		ArmourAmountMin int    `json:"armour_amount_min"`
		ArmourAmountMax int    `json:"armour_amount_max"`
		IsLegendary     bool   `json:"is_legendary"`
		ImageURL        string `json:"imageURL"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	armour := &data.Armour{
		ID:              uuid.New().String(),
		Type:            input.Type,
		BaseName:        input.BaseName,
		MinLevel:        input.MinLevel,
		ArmourAmountMin: input.ArmourAmountMin,
		ArmourAmountMax: input.ArmourAmountMax,
		IsLegendary:     input.IsLegendary,
		ImageURL:        input.ImageURL,
	}

	err := app.models.Armours.Insert(armour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	c.IndentedJSON(http.StatusCreated, armour)
}
