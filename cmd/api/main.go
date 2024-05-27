package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/utils"
	"context"
)

type application struct {
	models data.Models
}

func main() {
	utils.LoadEnv()

	ctx := context.Background()

	db, err := utils.CreateDynamoDBClient(ctx)
	if err != nil {
		panic(constants.DynamoDBClientError.Error())
	}

	app := &application{
		models: data.NewModels(db, ctx),
	}

	app.setupRoutes()
}
