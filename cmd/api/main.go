package main

import (
	"browser-mmo-backend/internal/constants"
	"browser-mmo-backend/internal/helpers"
	"context"
)

type application struct {
	// models data.Models
}

func main() {
	ctx := context.Background()

	_, err := helpers.CreateDynamoDBClient(ctx)
	if err != nil {
		panic(constants.ErrDynamoDBClient.Error())
	}

	app := &application{}

	app.setupRoutes()
}
