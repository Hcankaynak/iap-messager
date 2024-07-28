package messages

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
)

type MessageRepository struct {
	db *gorm.DB
}

func (m *MessageRepository) FindSentMessages() ([]Message, error) {
	var messages []Message
	err := m.db.Where("sending_status = ?", true).Find(&messages).Error
	return messages, err
}

func GenerateMessagesFromDummyData() []Message {
	jsonFile, err := os.Open("data/dummy_data.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var messages []Message
	if err := json.Unmarshal(byteValue, &messages); err != nil {
		log.Fatalf("failed to unmarshal JSON data: %v", err)
	}

	for _, message := range messages {
		log.Println(message)
	}
	return messages
}
