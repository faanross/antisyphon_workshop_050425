package websocket

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var WebSocketPort = 8080

var upgrader = websocket.Upgrader{
	// Allow connections from any origin for development
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Global connection registry to track all active connections
var (
	connections   = make(map[*websocket.Conn]bool)
	connMutex     = sync.Mutex{}
	messageBuffer = make(chan []byte, 100) // Buffer for messages to be sent
)

// WebSocketServer represents a simple WebSocket server
type WebSocketServer struct {
	port int
}

// NewWebSocketServer creates a new WebSocket server
func NewWebSocketServer(port int) *WebSocketServer {
	return &WebSocketServer{
		port: port,
	}
}

// Start begins the WebSocket server
func (s *WebSocketServer) Start() error {
	// Start the message broadcaster in a goroutine
	go s.broadcastMessages()

	// Set up HTTP handler for the WebSocket endpoint
	http.HandleFunc("/ws", s.handleWebSocket)

	// Start the server
	addr := fmt.Sprintf(":%d", s.port)
	fmt.Printf("WebSocket server starting on %s\n", addr)

	// Start the HTTP server (this is a blocking call)
	return http.ListenAndServe(addr, nil)
}

// handleWebSocket handles WebSocket connections
func (s *WebSocketServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	// Register the new connection
	connMutex.Lock()
	connections[conn] = true
	connMutex.Unlock()

	// Log the new connection
	fmt.Println("New WebSocket connection established")

	// Send a welcome message
	err = conn.WriteMessage(websocket.TextMessage, []byte("Connected to Go WebSocket Server"))
	if err != nil {
		log.Printf("Error sending message: %v", err)
		return
	}
	// Handle incoming messages in a goroutine
	go handleConnection(conn)

}

// handleConnection manages a single WebSocket connection
func handleConnection(conn *websocket.Conn) {
	defer func() {
		removeConnection(conn)
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading message: %v", err)
			}
			break
		}
		log.Printf("Received message: %s", message)

		// Echo the message back
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			break
		}
	}
}

// removeConnection safely removes a connection from the registry
func removeConnection(conn *websocket.Conn) {
	connMutex.Lock()
	delete(connections, conn)
	connMutex.Unlock()
	log.Println("WebSocket connection closed")
}

// broadcastMessages sends messages from the buffer to all connected clients
func (s *WebSocketServer) broadcastMessages() {
	for message := range messageBuffer {
		connMutex.Lock()
		for conn := range connections {
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error broadcasting message: %v", err)
				conn.Close()
				delete(connections, conn)
			}
		}
		connMutex.Unlock()
	}
}

// BroadcastMessage sends a message to all connected WebSocket clients
func BroadcastMessage(message []byte) {
	messageBuffer <- message
}

func StartWebSocketServer() {
	// Start WebSocket server in a separate goroutine
	wsServer := NewWebSocketServer(WebSocketPort)
	fmt.Printf("Starting WebSocket server on port %d...\n", WebSocketPort)
	go func() {
		err := wsServer.Start()
		if err != nil {
			log.Fatalf("WebSocket server error: %v", err)
		}
	}()

	// Give the WebSocket server a moment to start
	time.Sleep(100 * time.Millisecond)
	fmt.Println("WebSocket server is running. You can now connect from the web UI.")

}
