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

	// Send to the broadcaster
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

// GetAllListenersFromService gets all listeners from the service
func GetAllListenersFromService() ([]model.ListenerInfo, error) {
	// Get the service reference
	service := model.GetServiceProvider()
	if service == nil {
		return nil, fmt.Errorf("listener service not initialized")
	}

	// Get all listeners
	return service.GetAllListeners(), nil
}

// CreateListenerStatusMessage creates a listener status message without sending it
func CreateListenerStatusMessage(listeners []model.ListenerInfo) ([]byte, error) {
	message := WebSocketMessage{
		Type:    TypeListenerStatus,
		Payload: listeners,
		Time:    time.Now(),
	}

	return json.Marshal(message)
}
