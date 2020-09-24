package api

// ServerConfig defines various configurations.
type ServerConfig struct {
	// RabbitMQ connection settings.
	RabbitMQ struct {
		User     string
		Password string
		IP       string
		Port     int
	}

	// RestAPI defines URL to serve.
	RestAPI struct {
		IP   string
		Port int
	}
}
