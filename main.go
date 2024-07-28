package main

import (
	"github.com/Hcankaynak/iap-messager/configs"
	"github.com/Hcankaynak/iap-messager/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Welcome iap-messager!! ")

	configs.LoadEnv()

	db := database.ConnectPostgres(configs.LoadPostgres().GetDSN())
	database.Migrate(db, &User{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
