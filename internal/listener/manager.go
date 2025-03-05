package listener

import (
	"fmt"
	"sync"
)

// Manager keeps track of all active listeners
type Manager struct {
	listeners map[string]*Listener
	mu        sync.RWMutex
}

// NewManager creates a new listener manager
func NewManager() *Manager {
	return &Manager{
		listeners: make(map[string]*Listener),
	}
}

// AddListener adds a listener to the manager
func (m *Manager) AddListener(l *Listener) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.listeners[l.ID] = l
}

// PrintStatus prints the current status of all managed listeners
func (m *Manager) PrintStatus() {
	m.mu.RLock()
	defer m.mu.RUnlock()

	fmt.Printf("|UPDATE| Total Listeners: %d\n", len(m.listeners))

	if len(m.listeners) > 0 {
		fmt.Printf("|UPDATE| ID(s):")
		for id := range m.listeners {
			fmt.Printf(" %s |", id)
		}
	}
	fmt.Printf("\n")
}
