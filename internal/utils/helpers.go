package utils

import (
	"browser-mmo-backend/internal/constants"
	"context"
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
