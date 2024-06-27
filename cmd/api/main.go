package main

import (
	"browser-mmo-backend/constants"
	"browser-mmo-backend/data"
	"browser-mmo-backend/utils"
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
		panic(constants.DynamoDBClientError)
	}

	app := &application{
		models: data.NewModels(db, ctx),
	}

	app.setupRoutes()
}
