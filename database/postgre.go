package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectPostgres(dsn string) (db *gorm.DB) {
	log.Println("Connecting to Postgres...")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to Postgres!")
	return db
}

func Migrate(db *gorm.DB, models ...interface{}) {
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

}
