package utils

import (
	"browser-mmo-backend/internal/constants"
	"os"

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

func GetJWTPrivateKey() []byte {
	jwtPrivateKey := os.Getenv("JWT_PRIVATE_KEY")

	if jwtPrivateKey == "" {
		panic(constants.JWTPrivateKeyError)
	}

	return []byte(jwtPrivateKey)
}
