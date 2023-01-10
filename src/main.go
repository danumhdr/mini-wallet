package main

import (
	"context"
	"log"
	"mini-wallet/src/drivers/mongo"
	"mini-wallet/src/interfaces/http"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	db, errDB := mongo.Connect(ctx)
	if errDB != nil {
		log.Fatal("Error DB Connection")
	}

	api := http.New(&http.Http{
		DB: db,
	})
	api.Launch()
}
