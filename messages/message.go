package messages

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Content       string `json:"content"`
	PhoneNumber   string `json:"phone_number"`
	SendingStatus bool   `json:"sending_status"`
}
