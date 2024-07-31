package main

import (
	"github.com/Hcankaynak/iap-messager/configs"
	"github.com/Hcankaynak/iap-messager/database"
	"github.com/Hcankaynak/iap-messager/handlers"
	"log"
	"time"
)

// @title IAP Messager API
// @version 1.0
// @description This is a sample server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Welcome iap-messager!! ")
	time.Sleep(5 * time.Second)
	configs.LoadEnv()
	postgresDB := database.PostgresDB{Dsn: configs.LoadPostgres().GetDSN()}
	postgresDB.ConnectPostgres()
	postgresDB.Migrate()
	redis := database.NewRedisManager(configs.LoadRedisConnectionDataFromEnv())
	redis.ConnectRedis()
	redis.CreateEmptyListForInProgress()

	handlers.InitHandlers(postgresDB.GetDB(), &redis)
}
