package database

import (
	"github.com/Hcankaynak/iap-messager/messages"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

/*
PostgresDB struct
This struct is used to connect to the Postgres database.
*/
type PostgresDB struct {
	db  *gorm.DB
	Dsn string
}

/*
ConnectPostgres function
Connecting to the Postgres database.
*/
func (p *PostgresDB) ConnectPostgres() {
	log.Println("Connecting to Postgres...")

	db, err := gorm.Open(postgres.Open(p.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected to Postgres!")
	p.db = db
}

/*
Migrate function
Migrating the database.
Delete all data from the messages table and insert dummy data.
*/
func (p *PostgresDB) Migrate() {
	// make sure that the messages table is created
	err := p.db.AutoMigrate(&messages.Message{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// delete all data from messages table
	if err := p.db.Exec("DELETE FROM messages").Error; err != nil {
		log.Fatal("Failed to delete data from messages table:", err)
	}

	// insert dummy data
	messagesData := messages.GenerateMessagesFromDummyData()
	for _, message := range messagesData {
		if err := p.db.Create(&message).Error; err != nil {
			log.Fatalf("Failed to insert message: %v", err)
		}
	}
}

func (p *PostgresDB) GetDB() (db *gorm.DB) {
	return p.db
}
