package main

import (
	"github.com/Hcankaynak/iap-messager/configs"
	"github.com/Hcankaynak/iap-messager/database"
	"github.com/Hcankaynak/iap-messager/handlers"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Welcome iap-messager!! ")

	configs.LoadEnv()
	postgresDB := database.PostgresDB{Dsn: configs.LoadPostgres().GetDSN()}
	postgresDB.ConnectPostgres()
	postgresDB.Migrate()

	handlers.InitHandlers(postgresDB.GetDB())
}
