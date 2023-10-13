package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) createWeaponHandler(c *gin.Context) {
	var input struct {
		BaseName           string `json:"base_name"`
		MinLevel           int    `json:"min_level"`
		DamageLowRangeMin  int    `json:"damage_low_min"`
		DamageLowRangeMax  int    `json:"damage_low_max"`
		DamageHighRangeMin int    `json:"damage_high_min"`
		DamageHighRangeMax int    `json:"damage_high_max"`
		IsLegendary        bool   `json:"is_legendary"`
		ImageURL           string `json:"imageURL"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError.Error())
		return
	}

	weapon := &data.Weapon{
		ID:                 uuid.New().String(),
		BaseName:           input.BaseName,
		MinLevel:           input.MinLevel,
		DamageLowRangeMin:  input.DamageLowRangeMin,
		DamageLowRangeMax:  input.DamageLowRangeMax,
		DamageHighRangeMin: input.DamageHighRangeMin,
		DamageHighRangeMax: input.DamageHighRangeMax,
		IsLegendary:        input.IsLegendary,
		ImageURL:           input.ImageURL,
	}

	err := app.models.Weapons.Insert(weapon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, weapon)
}
