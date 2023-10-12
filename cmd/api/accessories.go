package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) createAccessoryHandler(c *gin.Context) {
	var input struct {
		BaseName    string `json:"base_name"`
		Type        string `json:"type"`
		MinLevel    int    `json:"min_level"`
		IsLegendary bool   `json:"is_legendary"`
		ImageURL    string `json:"imageURL"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError.Error())
		return
	}

	accessory := &data.Accessory{
		ID:          uuid.New().String(),
		BaseName:    input.BaseName,
		MinLevel:    input.MinLevel,
		IsLegendary: input.IsLegendary,
		ImageURL:    input.ImageURL,
	}

	err := app.models.Accessories.Insert(accessory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, accessory)
}
