package db

// ServerConfig defines various configurations.
type ServerConfig struct {
	// RabbitMQ connection settings
	RabbitMQ struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		IP       string `yaml:"ip"`
		Port     int    `yaml:"port"`
	} `yaml:"RabbitMQ"`

	// MongoDB connection settings
	MongoDB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		IP       string `yaml:"ip"`
		Port     int    `yaml:"port"`
	} `yaml:"MongoDB"`
}
