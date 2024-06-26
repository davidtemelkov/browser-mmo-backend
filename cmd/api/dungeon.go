package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/fightsimulator"
	"browser-mmo-backend/internal/items"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) fightDungeonBossHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	boss, err := app.models.Bosses.Get(strconv.Itoa(user.Dungeon))
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userFighter := fightsimulator.NewFighterFromUser(*user)
	bossFighter := fightsimulator.NewFighterFromBoss(boss)

	fightLog, playerWon := fightsimulator.Simulate(userFighter, bossFighter)

	if playerWon {
		// TODO: Gain exp, gold and increment user.Dungeon by 1

		err = app.models.Users.LevelUp(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}

		// TODO: Add this logic for generating legendaries or getting the exp and gold reward
		item, err := items.GenerateItem(app.models.Weapons, app.models.Accessories, app.models.Shields, app.models.Armours, app.models.Users)
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}

		err = app.models.Users.AddItemToInventory(user, item)
		if err != nil {
			if err.Error() == constants.NoAvailableSlotError {
				c.JSON(http.StatusConflict, constants.NoAvailableSlotError)
				return
			}

			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"fightLog":        fightLog,
			"fightWon":        playerWon,
			"monsterName":     boss.Name,
			"monsterImageUrl": boss.ImageURL,
			"monsterLvl":      boss.Lvl,
			"monsterHealth":   boss.Constitution + 100, // TODO: Temporary
		})
}
