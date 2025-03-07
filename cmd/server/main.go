package main

import (
	"antisyphon_workshop_050425/internal/listener"
	"antisyphon_workshop_050425/internal/model"
	"antisyphon_workshop_050425/internal/websocket"
	"fmt"
	"time"
)

var serverPorts = []string{"7777", "8888", "9999"}

func main() {

	websocket.StartWebSocketServer()

	factory := listener.NewListenerFactory()
	manager := listener.NewManager()
	service := listener.NewService(factory, manager)

	// create global reference
	model.SetServiceProvider(service)

	for _, port := range serverPorts {
		time.Sleep(5 * time.Second)
		err := service.CreateAndStartListener(port)
		if err != nil {
			fmt.Printf("Error creating service: %v\n", err)
			continue
		}
	}

	select {}
}
