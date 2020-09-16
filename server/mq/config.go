package mq

import "fmt"

// SessionConfig contains configs to build Session.
type SessionConfig struct {
	// User for RabbitMQ authentification.
	User string

	// Password for RabbitMQ authentification.
	Password string

	// Addr to connect RabbitMQ. For example, "localhost:5672"
	Addr string

	// Exchanges to declare in RabbitMQ.
	Exchanges map[string]ExchangeConfig
}

// ExchangeConfig holds settings for RabbitMQ exchange.
type ExchangeConfig struct {
	// Name of exchange.
	Name string

	// Type of echange. Should be one of direct, fanout, topic, header.
	Type string

	// Queues contain queues to be bound to this exchange.
	Queues map[string]QueueConfig
}

// QueueConfig holds settings for RabbitMQ queue.
type QueueConfig struct {
	// Name of queue.
	Name string

	// RouteKey defines rountingKey of queue.
	RouteKey string
}

var (
	defaultConfig = &SessionConfig{
		User:     "guest",
		Password: "guest",
		Addr:     "localhost:5672",
		Exchanges: map[string]ExchangeConfig{
			"": {
				Name: "", // "amq.direct"
				Type: "", // "direct"
				Queues: map[string]QueueConfig{
					"hello": {
						Name:     "hello",
						RouteKey: "hello",
					},
				},
			},
		},
	}
)

// URL builds connector URL for RabbitMQ.
func (c *SessionConfig) URL() string {
	return fmt.Sprintf("amqp://%s:%s@%s", c.User, c.Password, c.Addr)
}