// websocket/messages.go
package websocket

import (
	"antisyphon_workshop_050425/internal/listener"
	"encoding/json"
	"fmt"
	"time"
)

// MessageType defines the type of message being sent
type MessageType string

const (
	// Message types
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
func SendListenerCreated(l *listener.Listener) error {
	info := ListenerInfo{
		ID:        l.ID,
		Port:      l.Port,
		CreatedAt: l.CreatedAt,
	}

	message := WebSocketMessage{
		Type:    TypeListenerCreated,
		Payload: info,
		Time:    time.Now(),
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshaling listener info: %w", err)
	}

	// Send to the broadcaster
	BroadcastMessage(messageBytes)
	return nil
}

// SendListenerStatus sends the current status of all listeners
func SendListenerStatus(listeners map[string]*listener.Listener) error {
	// Convert the map to a slice of ListenerInfo
	listenerInfos := make([]ListenerInfo, 0, len(listeners))
	for _, l := range listeners {
		listenerInfos = append(listenerInfos, ListenerInfo{
			ID:        l.ID,
			Port:      l.Port,
			CreatedAt: l.CreatedAt,
		})
	}

	message := WebSocketMessage{
		Type:    TypeListenerStatus,
		Payload: listenerInfos,
		Time:    time.Now(),
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshaling listener status: %w", err)
	}

	BroadcastMessage(messageBytes)
	return nil
}
