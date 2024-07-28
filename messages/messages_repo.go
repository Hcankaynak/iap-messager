package messages

import (
	"encoding/json"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
)

/*
MessageRepository struct
This struct is used to interact with the messages table.
*/
type MessageRepository struct {
	DB *gorm.DB
}

/*
FindSentMessages function
This function is used to find sent messages.
*/
func (m *MessageRepository) FindSentMessages() ([]Message, error) {
	var messages []Message
	err := m.DB.Where("sending_status = ?", true).Find(&messages).Error
	return messages, err
}

/*
GenerateMessagesFromDummyData function
This function is used to generate messages from dummy data.
*/
func GenerateMessagesFromDummyData() []Message {
	// opening dummy json file
	jsonFile, err := os.Open("data/dummy_data.json")
	if err != nil {
		panic(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}(jsonFile)

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	// unmarshalling JSON data
	var messages []Message
	if err := json.Unmarshal(byteValue, &messages); err != nil {
		log.Fatalf("failed to unmarshal JSON data: %v", err)
	}

	return messages
}
