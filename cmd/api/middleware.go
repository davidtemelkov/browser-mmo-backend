package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/utils"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (app *application) authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusForbidden, constants.MissingAuthorizationHeaderError)
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, constants.InvalidAuthorizationHeaderFormatError)
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return utils.GetJWTPrivateKey(), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, constants.InvalidTokenError)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, constants.InvalidTokenClaimsError)
			c.Abort()
			return
		}

		userEmail, exists := claims["email"].(string)
		if !exists {
			c.JSON(http.StatusBadRequest, constants.BadRequestError)
			c.Abort()
			return
		}

		user, err := app.models.Users.Get(userEmail)
		if err != nil {
			if errors.Is(err, errors.New(constants.UserNotFoundError)) {
				c.JSON(http.StatusNotFound, constants.UserNotFoundError)
				c.Abort()
				return
			}

			c.JSON(http.StatusInternalServerError, constants.InternalServerError)
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
