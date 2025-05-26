package api

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) setWorkHandler(c *gin.Context) {
	var input struct {
		Hours  int `json:"hours"`
		Reward int `json:"reward"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, constants.InvalidJSONFormatError)
		return
	}

	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	err := app.Models.Work.Set(user, input.Hours, input.Reward)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, "Work started")
}

func (app *Application) cancelWorkHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	err := app.Models.Work.Cancel(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, "Work cancelled")
}

func (app *Application) collectWorkRewardsHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	err := app.Models.Work.CollectRewards(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, "Rewards collected")
}
