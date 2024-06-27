package main

import (
	"browser-mmo-backend/constants"
	"browser-mmo-backend/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) createArmourHandler(c *gin.Context) {
	var input struct {
		BaseName    string `json:"base_name"`
		WhatItem    string `json:"whatItem"`
		MinLevel    int    `json:"min_level"`
		IsLegendary bool   `json:"is_legendary"`
		ImageURL    string `json:"imageURL"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	armour := &data.Armour{
		ID:          uuid.New().String(),
		WhatItem:    input.WhatItem,
		BaseName:    input.BaseName,
		MinLevel:    input.MinLevel,
		IsLegendary: input.IsLegendary,
		ImageURL:    input.ImageURL,
	}

	err := app.models.Armours.Insert(armour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	c.IndentedJSON(http.StatusCreated, armour)
}
