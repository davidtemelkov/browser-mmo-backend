package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/services"
	"browser-mmo-backend/internal/utils"
	"browser-mmo-backend/internal/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) registerUserHandler(c *gin.Context) {
	var input services.UserInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, constants.InvalidJSONFormatError.Error())
		return
	}

	user := services.GetInitialUser(input)

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

	jwt, err := utils.CreateJWT(user.Name, user.Email)
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

	jwt, err := utils.CreateJWT(user.Name, user.Email)
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
		//TODO: This shouldn't matter, users should be able to view others' profiles
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
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
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
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
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
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
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
		c.JSON(http.StatusInternalServerError, constants.InternalServerError.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
