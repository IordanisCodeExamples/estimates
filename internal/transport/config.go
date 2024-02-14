package transport

// Config is the configuration for the transportlayer
type Config struct {
	// Headers is a list of headers that are allowed to be sent to the service
	Headers []string
	// Hosts is a list of hosts are allowed to connect to the service
	Hosts []string
	// Methods is a list of methods that are allowed to be sent to the service
	Methods []string
}

// GetConfig returns the configuration for the transportlayer
func GetConfig() *Config {
	return &Config{
		Headers: []string{"Content-Type", "Origin", "Accept", "*"},
		Hosts:   []string{"*"},
		Methods: []string{"GET", "POST", "OPTIONS"},
	}
}
