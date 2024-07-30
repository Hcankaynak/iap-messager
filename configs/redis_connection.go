package configs

import (
	"github.com/Hcankaynak/iap-messager/database"
	"os"
)

func LoadRedisConnectionDataFromEnv() database.RedisConnection {
	return database.RedisConnection{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}
}
