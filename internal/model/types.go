package model

import (
	"time"
)

// ListenerInfo represents information about an HTTP listener
// This is used for communication between packages and over WebSocket
type ListenerInfo struct {
	ID        string    `json:"id"`
	Port      string    `json:"port"`
	CreatedAt time.Time `json:"createdAt"`
	Status    string    `json:"status,omitempty"` // will use this later
}
