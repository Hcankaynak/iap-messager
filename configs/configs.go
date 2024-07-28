package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	// Set environment variables from .env file
	log.Println("Loading environment variables...")

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully loaded environment variables!")
}
