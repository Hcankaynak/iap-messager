package handlers

import (
	"github.com/Hcankaynak/iap-messager/docs"
	"github.com/Hcankaynak/iap-messager/messages"
	"github.com/Hcankaynak/iap-messager/sender"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

/*
MessageHandler struct
This struct is used to handle messages.
*/
type MessageHandler struct {
	messageRepo   *messages.MessageRepository
	messageSender *sender.MessageSender
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
	messageSender := sender.New()
	messageHandler := MessageHandler{messageRepo: &messageRepo, messageSender: &messageSender}
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/api/v1/sent-messages", messageHandler.getSentMessages)
	r.POST("/api/v1/automatic-message-sender", messageHandler.automaticMessageSender)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080
}

// Get Sent Messages godoc
// @Summary Get Sent Messages
// @Description Get Sent Messages
// @Tags Get Sent Messages
// @Produce  json
// @Success 200 {object} []messages.MessageDTO
// @Router /sent-messages [get]
func (m *MessageHandler) getSentMessages(c *gin.Context) {
	// search for sentMessages that SendingStatus is true.
	messageEntities, err := m.messageRepo.FindSentMessages()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch sentMessages"})
		return
	}
	messageDTOs := m.messageRepo.ConvertEntityToDTO(messageEntities)
	c.JSON(200,
		gin.H{
			"sentMessages": messageDTOs},
	)
}

// Start/Stop Automatic Message Sender godoc
// @Summary Start/Stop Automatic Message Sender
// @Description Start/Stop Automatic Message Sender If you send a request with start=true, the automatic message sender
// will start. If you send a request with start=false, the automatic message sender will stop.
// @Tags Start/Stop Automatic Message Sender
// @Accept  json
// @Produce  json
// @Param   messageSender body AutomaticMessageSender true "Automatic Message Sender Payload"
// @Success 200 {string} string "Response of start/stop automatic message sender"
// @Router /automatic-message-sender [post]
func (m *MessageHandler) automaticMessageSender(c *gin.Context) {
	// start or stop automatic message sender based on request.
	var request AutomaticMessageSender
	if err := c.Bind(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	var responseText string
	if request.Start == true && m.messageSender.GetStartValue() == true {
		responseText = "Automatic message sender already started!"
	} else if request.Start == true && m.messageSender.GetStartValue() == false {
		responseText = "Automatic message sender started successfully!"
		m.messageSender.StartAutomaticMessageSender()
	} else if request.Start == false && m.messageSender.GetStartValue() == false {
		responseText = "Automatic message sender already stopped!"
	} else if request.Start == false && m.messageSender.GetStartValue() == true {
		responseText = "Automatic message sender stopped successfully!"
		m.messageSender.StopAutomaticMessageSender()
	}
	c.JSON(200, responseText)
}
