package listener

import (
	"antisyphon_workshop_050425/internal/model"
	"antisyphon_workshop_050425/internal/router"
	"fmt"
	"github.com/go-chi/chi/v5"
	"math/rand"
	"net/http"
	"time"
)

// ListenerFactory creates and manages HTTP listeners
type ListenerFactory struct{}

// NewListenerFactory creates a new listener factory
func NewListenerFactory() *ListenerFactory {
	return &ListenerFactory{}
}

// Listener represents an HTTP server instance
type Listener struct {
	ID        string
	Port      string
	Router    *chi.Mux
	CreatedAt time.Time
}

// CreateListener generates a new listener with a random port and unique ID
func (f *ListenerFactory) CreateListener(port string) (*Listener, error) {
	// Generate a random ID (6 digits)
	id := fmt.Sprintf("listener_%06d", rand.Intn(1000000))

	r := chi.NewRouter()

	router.SetupRoutes(r)

	fmt.Printf("|CREATE| Listener %s serving on %s\n", id, port)

	return &Listener{
		ID:        id,
		Port:      port,
		Router:    r,
		CreatedAt: time.Now(),
	}, nil
}

func (l *Listener) Start() error {
	addr := fmt.Sprintf(":%s", l.Port)
	fmt.Printf("|START| Listener %s serving on %s\n\n", l.ID, addr)
	return http.ListenAndServe(addr, l.Router)
}

// ToInfo converts a Listener to a ListenerInfo for sharing with other packages
func (l *Listener) ToInfo() model.ListenerInfo {
	return model.ListenerInfo{
		ID:        l.ID,
		Port:      l.Port,
		CreatedAt: l.CreatedAt,
		Status:    "running",
	}
}
