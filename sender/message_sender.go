package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	start            bool
	ticker           *time.Ticker
	isInitialStarted bool
}

func New() MessageSender {
	return MessageSender{start: false, ticker: time.NewTicker(DURATION), isInitialStarted: false}
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
	log.Printf("Response status: %s", resp.Status)
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
					// send messages
					fmt.Println("Hello, World! ", time.Now())

					//m.sendMessage("Hello, World!", "+1234567890")
				}
			}
		}()
		m.isInitialStarted = true
	} else {
		log.Println("[DEBUG] Ticker is already started!")
	}
}
