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
	err := m.DB.Where("sent_status = ?", true).Find(&messages).Error
	return messages, err
}

func (m *MessageRepository) ConvertEntityToDTO(messages []Message) []MessageDTO {
	var messageDTOs []MessageDTO
	for _, message := range messages {
		messageDTOs = append(messageDTOs, MessageDTO{
			Content:     message.Content,
			PhoneNumber: message.PhoneNumber,
			SentStatus:  message.SentStatus,
		})
	}
	return messageDTOs

}

/*
GenerateMessagesFromDummyData function
This function is used to generate messages from dummy data.
*/
func GenerateMessagesFromDummyData() []Message {
	// opening dummy json file
	jsonFile, err := os.Open("/app/data/dummy_data.json")
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

/*
FindOldestTwoMessagesThatNotSent function
This function is used to find the oldest two messages that are not sent.
*/
func (m *MessageRepository) FindOldestTwoMessagesThatNotSent(excludedMessages []string) ([]Message, error) {
	var messages []Message
	req := m.DB.Where("sent_status = ?", false).Order("created_at ASC").Limit(2)
	if len(excludedMessages) > 0 {
		req.Where("id NOT IN ?", excludedMessages)
	}
	err := req.Find(&messages).Error
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return nil, err
	}

	return messages, nil
}

/*
SetMessageAsSent function
This function is used to set a message as sent.
*/
func (m *MessageRepository) SetMessageAsSent(message Message) {
	result := m.DB.Model(&message).Update("sent_status", true)
	if result.Error != nil {
		log.Fatalf("Failed to update message: %v", result.Error)
	}
}
