package helpers

import (
	"browser-mmo-backend/internal/constants"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetAWSAccessKey() string {
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")

	if awsAccessKey == "" {
		panic(constants.ErrAWSAccessKey.Error())
	}

	return awsAccessKey
}

func GetAWSSecretKey() string {
	awsSecretKey := os.Getenv("AWS_SECRET_KEY")

	if awsSecretKey == "" {
		panic(constants.ErrAWSSecretKey.Error())
	}

	return awsSecretKey
}

func CreateDynamoDBClient(ctx context.Context) (*dynamodb.Client, error) {
	// awsAccessKeyID := GetAWSAccessKey()
	// awsSecretAccessKey := GetAWSSecretKey()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(constants.AWSRegion))
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}
