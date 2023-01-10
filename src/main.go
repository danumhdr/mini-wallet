package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	//ctx := context.Background()
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
}
