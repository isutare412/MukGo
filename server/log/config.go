package log

// ServerConfig defines various configurations.
type ServerConfig struct {
	// RabbitMQ connection settings
	RabbitMQ struct {
		User     string
		Password string
		IP       string
		Port     int
	}
}
