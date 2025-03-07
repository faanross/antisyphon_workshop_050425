package listener

import (
	"antisyphon_workshop_050425/internal/model"
	"antisyphon_workshop_050425/internal/websocket"
	"fmt"
)

// Service coordinates listener operations
type Service struct {
	factory *ListenerFactory
	manager *Manager
}

// NewService creates a new listener service
func NewService(f *ListenerFactory, m *Manager) *Service {
	return &Service{
		factory: f,
		manager: m,
	}
}

// CreateAndStartListener creates a new listener, adds it to the manager, and starts it
func (s *Service) CreateAndStartListener(port string) error {
	// Create a new listener

	listener, err := s.factory.CreateListener(port)
	if err != nil {
		return err
	}

	// Add it to the manager
	s.manager.AddListener(listener)
	s.manager.PrintStatus()

	// Convert to ListenerInfo and notify the WebSocket clients
	listenerInfo := listener.ToInfo()
	err = websocket.SendListenerCreated(listenerInfo)
	if err != nil {
		fmt.Printf("Error sending listener creation notification: %v\n", err)
		// Continue even if notification fails
	}

	// Start the listener in a new goroutine
	go func(l *Listener) {

		err := listener.Start()
		if err != nil {
			fmt.Printf("Error starting listener %s: %v\n", l.ID, err)
		}
	}(listener)

	return nil
}

// GetAllListeners returns information about all managed listeners
func (s *Service) GetAllListeners() []model.ListenerInfo {
	return s.manager.GetAllListenersInfo()
}
