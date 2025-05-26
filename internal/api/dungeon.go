package api

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/fightsimulator"
	"browser-mmo-backend/internal/items"
	"browser-mmo-backend/internal/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *Application) getDungeonBossByPositionHandler(c *gin.Context) {
	postition := c.Param("position")
	positionInt, err := strconv.Atoi(postition)
	if err != nil || positionInt < 0 {
		c.JSON(http.StatusBadRequest, "Invalid boss position")
		return
	}

	boss, err := app.GameContent.GetBossByPosition(positionInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, boss)
}

func (app *Application) fightDungeonBossHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	boss, err := app.GameContent.GetBossByPosition(user.Dungeon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userFighter := fightsimulator.NewFighterFromUser(*user)
	bossFighter := fightsimulator.NewFighterFromBoss(boss)

	fightLog, playerWon := fightsimulator.Simulate(userFighter, bossFighter)

	if playerWon {
		err = app.Models.Users.CollectDungeonFightRewards(user, 0, 0) // Todo: change 0 0 to boss.exp boss.gold rewards
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}

		expForNextLvl := users.CalculateExpForLvlUp(user.Lvl)
		if user.EXP >= expForNextLvl {
			err = app.Models.Users.LevelUp(user, expForNextLvl)
			if err != nil {
				c.JSON(http.StatusInternalServerError, constants.InternalServerError)
				return
			}
		}

		// TODO: Either get legendary or big exp and gold reward
		item, err := items.GenerateItem(true, *user, app.GameContent.Weapons, app.GameContent.Accessories, app.GameContent.Shields, app.GameContent.Armours, app.Models.Users)
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}

		err = app.Models.Users.AddItemToInventory(user, item)
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
			"fightLog":      fightLog,
			"fightWon":      playerWon,
			"monsterId":     boss.ID,
			"monsterName":   boss.Name,
			"monsterLvl":    boss.Lvl,
			"monsterHealth": boss.Constitution + 100, // TODO: Temporary
		})
}
