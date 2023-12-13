package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) setWorkHandler(c *gin.Context) {
	var input struct {
		Hours  int `json:"hours"`
		Reward int `json:"reward"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, constants.InvalidJSONFormatError.Error())
		return
	}

	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	err := app.models.Work.Set(user, input.Hours, input.Reward)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "Work started")
}

func (app *application) cancelWorkHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	err := app.models.Work.Cancel(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "Work cancelled")
}

func (app *application) collectWorkRewardsHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	err := app.models.Work.CollectRewards(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "Rewards collected")
}
