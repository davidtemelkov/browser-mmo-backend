package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/helpers"
	"browser-mmo-backend/internal/validator"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
		Items: map[string]string{
			"Helmet":     "",
			"Chestplate": "",
			"Amulet":     "",
			"Gloves":     "",
			"Boots":      "",
			"Weapon":     "",
			"Shield":     "",
			"Ring":       "",
		},
		WeaponShop:    map[string]string{"1": "", "2": "", "3": "", "4": "", "5": "", "6": ""},                                                                                        //TODO: Fill this in with random items later
		MagicShop:     map[string]string{"1": "", "2": "", "3": "", "4": "", "5": "", "6": ""},                                                                                        //TODO: Fill this in with random items later
		Inventory:     map[string]string{"1": "", "2": "", "3": "", "4": "", "5": "", "6": "", "7": "", "8": "", "9": "", "10": "", "11": "", "12": "", "13": "", "14": "", "15": ""}, //TODO: Fill this in with random items later
		Mount:         "",
		MountImageURL: "",
		IsQuesting:    false,
		IsWorking:     false,
		CurrentQuests: map[string]data.GeneratedQuest{},
	}

	err := user.Password.Set(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
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
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	jwt, err := helpers.CreateJWT(user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
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

	claims, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	userEmail, exists := claims.(jwt.MapClaims)["email"].(string)
	if !exists {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	user, err := app.models.Users.Get(email)
	if err != nil {
		if errors.Is(err, constants.UserNotFoundError) {
			c.JSON(http.StatusNotFound, constants.UserNotFoundError.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if email != userEmail {
		c.JSON(http.StatusForbidden, constants.UserIsNotAuthorizedError.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (app *application) upgradeStrengthHandler(c *gin.Context) {
	//this should be moved so it can be reused, maybe in a middleware
	claims, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	userEmail, exists := claims.(jwt.MapClaims)["email"].(string)
	if !exists {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	user, err := app.models.Users.Get(userEmail)
	if err != nil {
		if errors.Is(err, constants.UserNotFoundError) {
			c.JSON(http.StatusNotFound, constants.UserNotFoundError.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//end

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
	//this should be moved so it can be reused, maybe in a middleware
	claims, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	userEmail, exists := claims.(jwt.MapClaims)["email"].(string)
	if !exists {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	user, err := app.models.Users.Get(userEmail)
	if err != nil {
		if errors.Is(err, constants.UserNotFoundError) {
			c.JSON(http.StatusNotFound, constants.UserNotFoundError.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//end

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
	//this should be moved so it can be reused, maybe in a middleware
	claims, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	userEmail, exists := claims.(jwt.MapClaims)["email"].(string)
	if !exists {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	user, err := app.models.Users.Get(userEmail)
	if err != nil {
		if errors.Is(err, constants.UserNotFoundError) {
			c.JSON(http.StatusNotFound, constants.UserNotFoundError.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//end

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
	//this should be moved so it can be reused, maybe in a middleware
	claims, ok := c.Get("user")
	if !ok {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	userEmail, exists := claims.(jwt.MapClaims)["email"].(string)
	if !exists {
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	user, err := app.models.Users.Get(userEmail)
	if err != nil {
		if errors.Is(err, constants.UserNotFoundError) {
			c.JSON(http.StatusNotFound, constants.UserNotFoundError.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	//end

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
