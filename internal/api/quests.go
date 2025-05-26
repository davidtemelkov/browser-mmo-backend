package api

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/fightsimulator"
	"browser-mmo-backend/internal/items"
	"browser-mmo-backend/internal/monsters"
	"browser-mmo-backend/internal/quests"
	"browser-mmo-backend/internal/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) generateQuestsHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	generatedQuests, err := quests.GenerateQuestsForUser(app.GameContent.Quests)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	err = app.Models.Quests.SetQuests(user.Email, generatedQuests)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusCreated, generatedQuests)
}

func (app *Application) setCurrentQuestHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	var currentQuestMap map[string]quests.GeneratedQuest
	if err := c.BindJSON(&currentQuestMap); err != nil {
		c.JSON(http.StatusBadRequest, constants.InvalidJSONFormatError)
		return
	}

	//TODO: This could be in a single request
	err := app.Models.Quests.SetCurrentQuest(user.Email, currentQuestMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	err = app.Models.Quests.SetQuests(user.Email, []quests.GeneratedQuest{
		{Name: "Empty Quest 0", ID: "", Time: "", EXP: "0", Gold: "0"},
		{Name: "Empty Quest 1", ID: "", Time: "", EXP: "0", Gold: "0"},
		{Name: "Empty Quest 2", ID: "", Time: "", EXP: "0", Gold: "0"},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}
	//end

	c.IndentedJSON(http.StatusOK, currentQuestMap)
}

func (app *Application) cancelCurrentQuestHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	err := app.Models.Quests.CancelCurrentQuest(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, "quest cancelled")
}

func (app *Application) collectCurrentQuestRewardsHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	generatedMonster, err := monsters.GenerateMonster(app.GameContent.Monsters, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userFighter := fightsimulator.NewFighterFromUser(*user)
	monsterFighter := fightsimulator.NewFighterFromMonster(generatedMonster)

	fightLog, playerWon := fightsimulator.Simulate(userFighter, monsterFighter)

	if playerWon {
		err = app.Models.Quests.CollectCurrentQuestRewards(user)
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

		// TODO: Make reward not always be item, if no item more exp and gold
		item, err := items.GenerateItem(false, *user, app.GameContent.Weapons, app.GameContent.Accessories, app.GameContent.Shields, app.GameContent.Armours, app.Models.Users)
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
	} else {
		err = app.Models.Quests.CancelCurrentQuest(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"fightLog":      fightLog,
			"fightWon":      playerWon,
			"monsterName":   generatedMonster.Name,
			"monsterId":     generatedMonster.ID,
			"monsterLvl":    generatedMonster.Lvl,
			"monsterHealth": monsterFighter.Health,
		})
}

func (app *Application) resetQuestsHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	err := app.Models.Quests.ResetQuests(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, "quests reset")
}
