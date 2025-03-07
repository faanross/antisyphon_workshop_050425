package model

// ServiceProvider defines the interface for accessing listener information
type ServiceProvider interface {
	GetAllListeners() []ListenerInfo
}

// Global reference to the current service instance
var currentServiceProvider ServiceProvider

// SetServiceProvider sets the global service reference
func SetServiceProvider(service ServiceProvider) {
	currentServiceProvider = service
}

// GetServiceProvider returns the global service reference
func GetServiceProvider() ServiceProvider {
	return currentServiceProvider
}
