package sender

import (
	"bytes"
	"encoding/json"
	"github.com/Hcankaynak/iap-messager/database"
	"github.com/Hcankaynak/iap-messager/messages"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
DURATION constant
This constant is used to set the duration of the ticker.
*/
const DURATION = 5 * time.Second

/*
MessageSender struct
This struct is used to send messages.
*/
type MessageSender struct {
	start             bool
	ticker            *time.Ticker
	isInitialStarted  bool
	redisManager      *database.RedisManager
	messageRepository *messages.MessageRepository
}

func New(redisManager *database.RedisManager, messageRepo *messages.MessageRepository) MessageSender {
	return MessageSender{start: false, ticker: time.NewTicker(DURATION), isInitialStarted: false, redisManager: redisManager, messageRepository: messageRepo}
}

func (m *MessageSender) sendMessage(content string, phoneNumber string) error {
	url := "https://webhook.site/299fcc61-3f22-4269-98f9-f557ea277a8f"

	// Create the request payload
	payload := map[string]string{
		"to":      phoneNumber,
		"content": content,
	}

	// Convert payload to JSON
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshalling payload: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-ins-auth-key", "INS.me1x9uMcyYGlhKKQVPoc.bO3j9aZwRTOcA2Ywo")

	// Create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("failed to close body	: %v", err)
		}
	}(resp.Body)

	// Log the response status
	log.Printf("Response Body: %s", resp.Body)
	return nil
}

/*
StartAutomaticMessageSender function
This function is used to start automatic message sender.
*/
func (m *MessageSender) StartAutomaticMessageSender() {
	m.start = true
	m.ticker.Reset(DURATION)
	m.startTicker()
}

/*
StopAutomaticMessageSender function
This function is used to stop automatic message sender.
*/
func (m *MessageSender) StopAutomaticMessageSender() {
	m.start = false
	m.ticker.Stop()

}

/*
GetStartValue function
This function is used to get the start value.
*/
func (m *MessageSender) GetStartValue() bool {
	return m.start
}

func (m *MessageSender) startTicker() {
	if !m.isInitialStarted {
		go func() {
			for range m.ticker.C {
				if m.start {
					m.findMessagesAndSend()
				}
			}
		}()
		m.isInitialStarted = true
	} else {
		log.Println("[DEBUG] Ticker is already started!")
	}
}

/*
findMessagesAndSend function
This function is used to find messages and send them.
*/
func (m *MessageSender) findMessagesAndSend() {
	messagesToBeSent, err := m.messageRepository.FindOldestTwoMessagesThatNotSent(m.redisManager.GetInProgressMessages().Items)
	if err != nil {
		log.Fatalf("Failed to find messages: %v", err)
	}
	for _, message := range messagesToBeSent {
		m.redisManager.AddInProgressMessage(strconv.Itoa(int(message.ID)))
		err := m.sendMessage(message.Content, message.PhoneNumber)
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
		m.redisManager.SetItem(strconv.Itoa(int(message.ID)), message.ID)
		m.messageRepository.SetMessageAsSent(message)
		m.redisManager.RemoveFromInProgressMessages(strconv.Itoa(int(message.ID)))
	}
}
