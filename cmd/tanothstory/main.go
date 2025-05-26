package main

import (
	"browser-mmo-backend/internal/api"
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/gamecontent"
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(constants.LoadingEnvFileError)
	}

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(constants.AWSRegion))
	if err != nil {
		panic(constants.DynamoDBClientError)
	}
	db := dynamodb.NewFromConfig(cfg)

	gamecontent, err := gamecontent.LoadGameContent(constants.GameContentDir)
	if err != nil {
		panic("Failed to load game content: " + err.Error())
	}

	app := &api.Application{
		Models:      data.NewModels(db, ctx),
		GameContent: *gamecontent,
	}

	app.SetupRoutes()
}
