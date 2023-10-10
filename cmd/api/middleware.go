package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (app *application) authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusForbidden, constants.MissingAuthorizationHeaderError.Error())
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, constants.InvalidAuthorizationHeaderFormatError.Error())
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return helpers.GetJWTPrivateKey(), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, constants.InvalidTokenError.Error())
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, constants.InvalidTokenClaimsError.Error())
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
