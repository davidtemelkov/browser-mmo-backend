package api

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/fightsimulator"
	"browser-mmo-backend/internal/items"
	"browser-mmo-backend/internal/users"
	"browser-mmo-backend/internal/utils"
	"browser-mmo-backend/internal/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) registerUserHandler(c *gin.Context) {
	var input users.UserInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	user := users.GetInitialUser(input)

	err := user.Password.Set(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	v := validator.New()
	if data.ValidateRegisterInput(v, user); !v.Valid() {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": v.Errors})
		return
	}

	err = app.Models.Users.Insert(user)
	if err != nil {
		if err.Error() == constants.DuplicateEmailError {
			c.JSON(http.StatusConflict, constants.DuplicateEmailError)
			return
		}
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	jwt, err := utils.CreateJWT(user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.IndentedJSON(http.StatusCreated, jwt)
}

func (app *Application) loginUserHandler(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError)
		return
	}

	v := validator.New()
	if data.ValidateLoginInput(v, input.Email, input.Password); !v.Valid() {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": v.Errors})
		return
	}

	user, err := app.Models.Users.Get(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.FailedLoginError)
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, constants.FailedLoginError)
		return
	}

	userLoggedIn, err := app.Models.Users.CanLoginUser(input.Password, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	if !userLoggedIn {
		c.JSON(http.StatusNotFound, constants.FailedLoginError)
		return
	}

	jwt, err := utils.CreateJWT(user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, jwt)
}

// TODO: Remove unnecessary pointers to user structs and all others
func (app *Application) getUserHandler(c *gin.Context) {
	email := c.Param("email")

	user, err := app.Models.Users.Get(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	users.CalculateTotalStats(user)

	c.JSON(http.StatusOK, user)
}

func (app *Application) upgradeStrengthHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	if user.Gold < user.Strength {
		c.JSON(http.StatusForbidden, "Not enough gold")
		return
	}

	result, err := app.Models.Users.UpgradeStrength(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (app *Application) upgradeDexterityHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	if user.Gold < user.Dexterity {
		c.JSON(http.StatusForbidden, "Not enough gold")
		return
	}

	result, err := app.Models.Users.UpgradeDexterity(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (app *Application) upgradeConstitutionHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	if user.Gold < user.Constitution {
		c.JSON(http.StatusForbidden, "Not enough gold")
		return
	}

	result, err := app.Models.Users.UpgradeConstitution(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (app *Application) upgradeIntelligenceHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	if user.Gold < user.Intelligence {
		c.JSON(http.StatusForbidden, "Not enough gold")
		return
	}

	result, err := app.Models.Users.UpgradeIntelligence(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// TODO: Make logic for only weaponShop items to be generated and added
// TODO: Add shopsLastGenerated in dynamo
func (app *Application) generateWeaponShop(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	var generatedItems []data.Item
	for i := 0; i < 6; i++ {
		item, err := items.GenerateItem(false, *user, app.GameContent.Weapons, app.GameContent.Accessories, app.GameContent.Shields, app.GameContent.Armours, app.Models.Users)
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}

		generatedItems = append(generatedItems, item)
	}

	err := app.Models.Users.GenerateWeaponShop(user, generatedItems)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "generated weapon shop")
}

// TODO: Make logic for only magicShop items to be generated and added
// TODO: Add shopsLastGenerated in dynamo
func (app *Application) generateMagicShop(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	var generatedItems []data.Item
	for i := 0; i < 6; i++ {
		item, err := items.GenerateItem(false, *user, app.GameContent.Weapons, app.GameContent.Accessories, app.GameContent.Shields, app.GameContent.Armours, app.Models.Users)
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}

		generatedItems = append(generatedItems, item)
	}

	err := app.Models.Users.GenerateMagicShop(user, generatedItems)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "generated weapon shop")
}

func (app *Application) equipItem(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	slotKey := c.Param("slotKey")

	err := app.Models.Users.EquipItem(user, slotKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "item equipped")
}

func (app *Application) buyMagicShopItem(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	slotKey := c.Param("slotKey")

	replacementItem, err := items.GenerateItem(false, *user, app.GameContent.Weapons, app.GameContent.Accessories, app.GameContent.Shields, app.GameContent.Armours, app.Models.Users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	err = app.Models.Users.BuyItem(user, slotKey, constants.MagicShopAttribute, replacementItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "item bought")
}

func (app *Application) buyWeaponShopItem(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	slotKey := c.Param("slotKey")

	replacementItem, err := items.GenerateItem(false, *user, app.GameContent.Weapons, app.GameContent.Accessories, app.GameContent.Shields, app.GameContent.Armours, app.Models.Users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	err = app.Models.Users.BuyItem(user, slotKey, constants.WeaponShopAttribute, replacementItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "item bought")
}

func (app *Application) sellItem(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	slotKey := c.Param("slotKey")

	err := app.Models.Users.SellItem(user, slotKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "item sold")
}

func (app *Application) fightPlayerHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	email := c.Param("email")
	enemy, err := app.Models.Users.Get(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userFighter := fightsimulator.NewFighterFromUser(*user)
	enemyFighter := fightsimulator.NewFighterFromUser(*enemy)

	fightLog, playerWon := fightsimulator.Simulate(userFighter, enemyFighter)

	if playerWon {
		err = app.Models.Users.CollectPlayerFightRewards(user, enemy)
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
	}

	c.IndentedJSON(http.StatusOK,
		gin.H{
			"fightLog":    fightLog,
			"fightWon":    playerWon,
			"userName":    enemy.Name,
			"userImageId": enemy.ImageID,
			"userLvl":     enemy.Lvl,
			"userHealth":  enemy.Constitution + 100, // TODO: Temporary
		})
}
