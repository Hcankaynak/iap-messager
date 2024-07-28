package configs

import (
	"github.com/joho/godotenv"
	"log"
)

/*
LoadEnv function
Loading environment variables from .env file.
*/
func LoadEnv() {
	log.Println("Loading environment variables...")

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully loaded environment variables!")
}
