package messages

import "gorm.io/gorm"

/*
Message struct
This struct is used to represent a message.
*/
type Message struct {
	gorm.Model
	Content     string `json:"content" gorm:"size:500"`
	PhoneNumber string `json:"phone_number" gorm:"size:30"`
	SentStatus  bool   `json:"sent_status"`
}

/*
MessageDTO struct
This struct is used to represent a message data transfer object.
*/
type MessageDTO struct {
	Content     string `json:"content" gorm:"size:500"`
	PhoneNumber string `json:"phone_number" gorm:"size:30"`
	SentStatus  bool   `json:"sent_status"`
}
