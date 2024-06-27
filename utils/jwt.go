package utils

import (
	"github.com/golang-jwt/jwt"
)

// JWT Claims constants
const (
	Name         = "name"
	EmailAddress = "email"
)

func CreateJWT(username, userEmail string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		Name:         username,
		EmailAddress: userEmail,
	})

	privateKey := GetJWTPrivateKey()

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
