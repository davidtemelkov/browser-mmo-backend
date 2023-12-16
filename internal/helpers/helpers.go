package helpers

import (
	"browser-mmo-backend/internal/constants"
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(constants.LoadingEnvFileError.Error())
	}
}

func GetJWTPrivateKey() []byte {
	jwtPrivateKey := os.Getenv("JWT_PRIVATE_KEY")

	if jwtPrivateKey == "" {
		panic(constants.JWTPrivateKeyError.Error())
	}

	return []byte(jwtPrivateKey)
}

func GetFirebaseUrl() string {
	firebaseUrl := os.Getenv("FIREBASE_URL")

	if firebaseUrl == "" {
		panic(constants.FirebaseURLKeyError.Error())
	}

	return firebaseUrl
}

func GetFirebaseBucketName() string {
	firebaseBucketName := os.Getenv("FIREBASE_BUCKET_NAME")

	if firebaseBucketName == "" {
		panic(constants.FirebaseBucketNameError.Error())
	}

	return firebaseBucketName
}

func CreateDynamoDBClient(ctx context.Context) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(constants.AWSRegion))
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}

func GetCurrentDate() string {
	currentTime := time.Now().UTC()

	return currentTime.Format(constants.TimeFormatJustDate)
}
