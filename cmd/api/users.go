package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/helpers"
	"browser-mmo-backend/internal/validator"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *application) registerUserHandler(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		ImageURL string `json:"imageURL"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError.Error())
		return
	}

	user := &data.User{
		Name:         input.Name,
		Email:        input.Email,
		CreatedOn:    time.Now().UTC().Format(constants.TimeFormat),
		ImageURL:     input.ImageURL,
		Level:        1,
		Gold:         100,
		EXP:          0,
		BigDPoints:   0,
		Strength:     1,
		Dexterity:    1,
		Constitution: 1,
		Intelligence: 1,
		Items: map[string]data.Item{
			"Helmet":     {ID: "0", Name: "Empty Helmet"},
			"Chestplate": {ID: "0", Name: "Empty Chestplate"},
			"Amulet":     {ID: "0", Name: "Empty Amulet"},
			"Gloves":     {ID: "0", Name: "Empty Gloves"},
			"Boots":      {ID: "0", Name: "Empty Boots"},
			"Weapon":     {ID: "0", Name: "Empty Weapon"},
			"Shield":     {ID: "0", Name: "Empty Shield"},
			"Ring":       {ID: "0", Name: "Empty Ring"},
		},
		WeaponShop: map[string]data.Item{
			"Item1": {ID: "0", Name: "Empty Item 1"},
			"Item2": {ID: "0", Name: "Empty Item 2"},
			"Item3": {ID: "0", Name: "Empty Item 3"},
			"Item4": {ID: "0", Name: "Empty Item 4"},
			"Item5": {ID: "0", Name: "Empty Item 5"},
			"Item6": {ID: "0", Name: "Empty Item 6"}}, //TODO: Fill this in with random items later
		MagicShop: map[string]data.Item{
			"Item1": {ID: "0", Name: "Empty Item 1"},
			"Item2": {ID: "0", Name: "Empty Item 2"},
			"Item3": {ID: "0", Name: "Empty Item 3"},
			"Item4": {ID: "0", Name: "Empty Item 4"},
			"Item5": {ID: "0", Name: "Empty Item 5"},
			"Item6": {ID: "0", Name: "Empty Item 6"}}, //TODO: Fill this in with random items later
		Inventory: map[string]data.Item{
			"Item1":  {ID: "0", Name: "Empty Item 1"},
			"Item2":  {ID: "0", Name: "Empty Item 2"},
			"Item3":  {ID: "0", Name: "Empty Item 3"},
			"Item4":  {ID: "0", Name: "Empty Item 4"},
			"Item5":  {ID: "0", Name: "Empty Item 5"},
			"Item6":  {ID: "0", Name: "Empty Item 6"},
			"Item7":  {ID: "0", Name: "Empty Item 7"},
			"Item8":  {ID: "0", Name: "Empty Item 8"},
			"Item9":  {ID: "0", Name: "Empty Item 9"},
			"Item10": {ID: "0", Name: "Empty Item 10"},
			"Item11": {ID: "0", Name: "Empty Item 11"},
			"Item12": {ID: "0", Name: "Empty Item 12"},
			"Item13": {ID: "0", Name: "Empty Item 13"},
			"Item14": {ID: "0", Name: "Empty Item 14"},
			"Item15": {ID: "0", Name: "Empty Item 15"}}, //TODO: Fill this in with random items later
		Mount:         "",
		MountImageURL: "",
		IsQuesting:    false,
		IsWorking:     false,
		CurrentQuests: map[string]data.GeneratedQuest{},
	}

	err := user.Password.Set(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	v := validator.New()
	if data.ValidateRegisterInput(v, user); !v.Valid() {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": v.Errors})
		return
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		if err == constants.DuplicateEmailError {
			c.JSON(http.StatusConflict, constants.DuplicateEmailError.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	jwt, err := helpers.CreateJWT(user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError.Error())
		return
	}

	v := validator.New()
	if data.ValidateLoginInput(v, input.Email, input.Password); !v.Valid() {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": v.Errors})
		return
	}

	user, err := app.models.Users.Get(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.FailedLoginError.Error())
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, constants.FailedLoginError.Error())
		return
	}

	userLoggedIn, err := app.models.Users.CanLoginUser(input.Password, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	if !userLoggedIn {
		c.JSON(http.StatusNotFound, constants.FailedLoginError.Error())
		return
	}

	jwt, err := helpers.CreateJWT(user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	c.JSON(http.StatusOK, jwt)
}

func (app *application) getUserHandler(c *gin.Context) {
	email := c.Param("email")

	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	if email != user.Email {
		c.JSON(http.StatusForbidden, constants.UserIsNotAuthorizedError.Error())
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
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusInternalServerError, err.Error())
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
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func (app *application) upgradeIntelligenceHandler(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*data.User)

	if user.Gold < user.Strength {
		c.JSON(http.StatusForbidden, "Not enough gold")
		return
	}

	result, err := app.models.Users.UpgradeIntelligence(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
