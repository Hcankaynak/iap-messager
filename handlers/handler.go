package handlers

import (
	"github.com/Hcankaynak/iap-messager/messages"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
MessageHandler struct
This struct is used to handle messages.
*/
type MessageHandler struct {
	messageRepo *messages.MessageRepository
}

/*
AutomaticMessageSender struct
This struct is used to read request of start or stop automatic message sender.
*/
type AutomaticMessageSender struct {
	Start bool `json:"start"`
}

/*
InitHandlers function
This function is used to initialize handlers.
*/
func InitHandlers(postgresDB *gorm.DB) {
	messageRepo := messages.MessageRepository{DB: postgresDB}
	messageHandler := MessageHandler{messageRepo: &messageRepo}
	r := gin.Default()
	r.GET("/sent-messages", messageHandler.getSentMessages)
	r.POST("/automatic-message-sender", messageHandler.automaticMessageSender)
	r.Run() // listen and serve on 0.0.0.0:8080
}

/*
getSentMessages function
This function is used to get sent messages.
*/
func (m MessageHandler) getSentMessages(c *gin.Context) {
	// search for sentMessages that SendingStatus is true.
	sentMessages, err := m.messageRepo.FindSentMessages()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch sentMessages"})
		return
	}
	c.JSON(200,
		gin.H{
			"sentMessages": sentMessages},
	)
}

/*
automaticMessageSender function
This function is used to start or stop automatic message sender.
*/
func (m MessageHandler) automaticMessageSender(c *gin.Context) {
	// start or stop automatic message sender based on request.
	var request AutomaticMessageSender
	if err := c.Bind(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// start or stop based on request.

	c.JSON(200, "bingos")
}
