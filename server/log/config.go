package log

// ServerConfig defines various configurations.
type ServerConfig struct {
	// RabbitMQ connection settings
	RabbitMQ struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		IP       string `yaml:"ip"`
		Port     int    `yaml:"port"`
	} `yaml:"RabbitMQ"`
}
