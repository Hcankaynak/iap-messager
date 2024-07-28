package messages

import "gorm.io/gorm"

/*
Message struct
This struct is used to represent a message.
*/
type Message struct {
	gorm.Model
	Content       string `json:"content"`
	PhoneNumber   string `json:"phone_number"`
	SendingStatus bool   `json:"sending_status"`
}
