package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/helpers"
	"context"
)

type application struct {
	models data.Models
}

func main() {
	helpers.LoadEnv()

	ctx := context.Background()

	db, err := helpers.CreateDynamoDBClient(ctx)
	if err != nil {
		panic(constants.DynamoDBClientError.Error())
	}

	app := &application{
		models: data.NewModels(db, ctx),
	}

	app.setupRoutes()
}
