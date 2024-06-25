package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/fightsimulator"
	"browser-mmo-backend/internal/items"
	"browser-mmo-backend/internal/monsters"
	"browser-mmo-backend/internal/quests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) createQuestHandler(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		ImageURL string `json:"imageURL"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	quest := &data.Quest{
		ID:       uuid.New().String(),
		Name:     input.Name,
		ImageURL: input.ImageURL,
	}

	err := app.models.Quests.Insert(quest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	c.IndentedJSON(http.StatusCreated, quest)
}

func (app *application) generateQuestsHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	allQuests, err := app.models.Quests.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	generatedQuests, err := quests.GenerateQuestsForUser(allQuests)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	err = app.models.Quests.SetQuests(user.Email, generatedQuests)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusCreated, generatedQuests)
}

func (app *application) setCurrentQuestHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	var currentQuestMap map[string]data.GeneratedQuest
	if err := c.BindJSON(&currentQuestMap); err != nil {
		c.JSON(http.StatusBadRequest, constants.InvalidJSONFormatError)
		return
	}

	//TODO: This could be in a single request
	err := app.models.Quests.SetCurrentQuest(user.Email, currentQuestMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	err = app.models.Quests.SetQuests(user.Email, []data.GeneratedQuest{
		{Name: "Empty Quest 0", ImageURL: "", Time: "", EXP: "0", Gold: "0"},
		{Name: "Empty Quest 1", ImageURL: "", Time: "", EXP: "0", Gold: "0"},
		{Name: "Empty Quest 2", ImageURL: "", Time: "", EXP: "0", Gold: "0"},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}
	//end

	c.IndentedJSON(http.StatusOK, currentQuestMap)
}

func (app *application) cancelCurrentQuestHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	err := app.models.Quests.CancelCurrentQuest(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, "quest cancelled")
}

func (app *application) collectCurrentQuestRewardsHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	generatedMonster, err := monsters.GenerateMonster(app.models.Monsters, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userFighter := fightsimulator.NewFighterFromUser(*user)
	monsterFighter := fightsimulator.NewFighterFromMonster(generatedMonster)

	fightLog, playerWon := fightsimulator.Simulate(userFighter, monsterFighter)

	if playerWon {
		err = app.models.Users.LevelUp(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}

		err = app.models.Quests.CollectCurrentQuestRewards(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}

		// TODO: Add this logic when generating quests so user can see what item
		// they will get, also make quest reward a little worse if there is an item
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
	} else {
		err = app.models.Quests.CancelCurrentQuest(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"fightLog":        fightLog,
			"fightWon":        playerWon,
			"monsterName":     generatedMonster.Name,
			"monsterImageUrl": generatedMonster.ImageURL,
			"monsterLvl":      generatedMonster.Lvl,
			"monsterHealth":   monsterFighter.Health,
		})
}

func (app *application) resetQuestsHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	err := app.models.Quests.ResetQuests(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, "quests reset")
}
