package main

import (
	"browser-mmo-backend/constants"
	"browser-mmo-backend/data"
	"browser-mmo-backend/items"
	"browser-mmo-backend/users"
	"browser-mmo-backend/utils"
	"browser-mmo-backend/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) registerUserHandler(c *gin.Context) {
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

	err = app.models.Users.Insert(user)
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

func (app *application) loginUserHandler(c *gin.Context) {
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

	user, err := app.models.Users.Get(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.FailedLoginError)
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, constants.FailedLoginError)
		return
	}

	userLoggedIn, err := app.models.Users.CanLoginUser(input.Password, user)
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
func (app *application) getUserHandler(c *gin.Context) {
	email := c.Param("email")

	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	if email != user.Email {
		//TODO: This shouldn't matter, users should be able to view others' profiles
		c.JSON(http.StatusForbidden, constants.UserIsNotAuthorizedError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (app *application) upgradeStrengthHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	if user.Gold < user.Strength {
		c.JSON(http.StatusForbidden, "Not enough gold")
		return
	}

	result, err := app.models.Users.UpgradeStrength(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (app *application) upgradeDexterityHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	if user.Gold < user.Dexterity {
		c.JSON(http.StatusForbidden, "Not enough gold")
		return
	}

	result, err := app.models.Users.UpgradeDexterity(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (app *application) upgradeConstitutionHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	if user.Gold < user.Constitution {
		c.JSON(http.StatusForbidden, "Not enough gold")
		return
	}

	result, err := app.models.Users.UpgradeConstitution(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (app *application) upgradeIntelligenceHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	if user.Gold < user.Intelligence {
		c.JSON(http.StatusForbidden, "Not enough gold")
		return
	}

	result, err := app.models.Users.UpgradeIntelligence(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// TODO: Make logic for only weaponShop items to be generated and added
func (app *application) generateWeaponShop(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	var generatedItems []data.Item
	for i := 0; i < 6; i++ {
		item, err := items.GenerateItem(false, *user, app.models.Weapons, app.models.Accessories, app.models.Shields, app.models.Armours, app.models.Users)
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}

		generatedItems = append(generatedItems, item)
	}

	err := app.models.Users.GenerateWeaponShop(user, generatedItems)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "generated weapon shop")
}

// TODO: Make logic for only magicShop items to be generated and added
func (app *application) generateMagicShop(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	var generatedItems []data.Item
	for i := 0; i < 6; i++ {
		item, err := items.GenerateItem(false, *user, app.models.Weapons, app.models.Accessories, app.models.Shields, app.models.Armours, app.models.Users)
		if err != nil {
			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			return
		}

		generatedItems = append(generatedItems, item)
	}

	err := app.models.Users.GenerateMagicShop(user, generatedItems)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "generated weapon shop")
}

func (app *application) equipItem(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	slotKey := c.Param("slotKey")

	err := app.models.Users.EquipItem(user, slotKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "item equipped")
}

func (app *application) buyMagicShopItem(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	slotKey := c.Param("slotKey")

	replacementItem, err := items.GenerateItem(false, *user, app.models.Weapons, app.models.Accessories, app.models.Shields, app.models.Armours, app.models.Users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	err = app.models.Users.BuyItem(user, slotKey, constants.MagicShopAttribute, replacementItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "item bought")
}

func (app *application) buyWeaponShopItem(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	slotKey := c.Param("slotKey")

	replacementItem, err := items.GenerateItem(false, *user, app.models.Weapons, app.models.Accessories, app.models.Shields, app.models.Armours, app.models.Users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	err = app.models.Users.BuyItem(user, slotKey, constants.WeaponShopAttribute, replacementItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "item bought")
}

func (app *application) sellItem(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	slotKey := c.Param("slotKey")

	err := app.models.Users.SellItem(user, slotKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, "item sold")
}
