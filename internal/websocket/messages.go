package websocket

import (
	"antisyphon_workshop_050425/internal/model"
	"encoding/json"
	"fmt"
	"time"
)

// MessageType defines the type of message being sent
type MessageType string

const (
	TypeListenerCreated MessageType = "listener_created"
	TypeListenerStatus  MessageType = "listener_status"
)

// WebSocketMessage is the structure for messages sent over the WebSocket
type WebSocketMessage struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload"`
	Time    time.Time   `json:"time"`
}

// ListenerInfo represents the listener data sent to the frontend
type ListenerInfo struct {
	ID        string    `json:"id"`
	Port      string    `json:"port"`
	CreatedAt time.Time `json:"createdAt"`
}

// SendListenerCreated sends a notification that a new listener was created
func SendListenerCreated(info model.ListenerInfo) error {
	fmt.Printf("Sending listener created notification for %s\n", info.ID)

	message := WebSocketMessage{
		Type:    TypeListenerCreated,
		Payload: info,
		Time:    time.Now(),
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshaling listener info: %w", err)
	}

	fmt.Printf("Marshaled message: %s\n", string(messageBytes))

	// Try to send the message, with retries if no clients are connected
	maxRetries := 5
	for attempt := 0; attempt < maxRetries; attempt++ {
		if IsClientConnected() {
			BroadcastMessage(messageBytes)
			return nil
		}

		fmt.Printf("No clients connected, retrying in 1 second (attempt %d/%d)\n",
			attempt+1, maxRetries)
		time.Sleep(1 * time.Second)
	}

	// Send it anyway, it will be queued in the message buffer
	BroadcastMessage(messageBytes)
	return nil
}

// SendListenerStatus sends the current status of all listeners
func SendListenerStatus(listeners []model.ListenerInfo) error {
	message := WebSocketMessage{
		Type:    TypeListenerStatus,
		Payload: listeners,
		Time:    time.Now(),
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshaling listener status: %w", err)
	}

	BroadcastMessage(messageBytes)
	return nil
}
