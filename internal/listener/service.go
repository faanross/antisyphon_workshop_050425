package listener

import (
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

	// Start the listener in a new goroutine
	go func(l *Listener) {

		err := listener.Start()
		if err != nil {
			fmt.Printf("Error starting listener %s: %v\n", l.ID, err)
		}
	}(listener)

	return nil
}
