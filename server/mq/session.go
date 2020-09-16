package mq

import (
	"fmt"
	"sync"
	"time"

	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/console"
	"github.com/streadway/amqp"
)

// Session keeps track of contructed connection and channel between server
// and message queue.
type Session struct {
	name   string
	config *SessionConfig

	conn *amqp.Connection
	ch   *amqp.Channel
	mu   sync.Mutex

	connNotifier *server.Broadcaster
}

var (
	sessionPool = make(map[string]*Session)
	mu          sync.Mutex
)

// NewSession creates session with ExchangeConfig. If config is nil, default
// is used.
func NewSession(name string, config *SessionConfig) *Session {
	if config == nil {
		config = defaultConfig
	}

	mu.Lock()
	defer mu.Unlock()

	// session already exists
	if s, ok := sessionPool[name]; ok {
		return s
	}

	s := &Session{
		name:         name,
		config:       config,
		connNotifier: server.NewBroadcaster(),
	}
	sessionPool[name] = s
	return s
}

// GetSession returns Session with name. Returns nil if the session does not
// exists.
func GetSession(name string) *Session {
	mu.Lock()
	defer mu.Unlock()
	return sessionPool[name]
}

// Connect establishes connection, channel, exchange, queue of RabbitMQ.
func (s *Session) Connect() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// establish connection
	conn, err := amqp.Dial(s.config.URL())
	if err != nil {
		return fmt.Errorf("on Session.Connect: %v", err)
	}
	s.conn = conn

	// open channel from connection
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("on Session.Connect: %v", err)
	}
	s.ch = ch

	for _, exchange := range s.config.Exchanges {
		// declare exchanges if not exists
		err := ch.ExchangeDeclare(
			exchange.Name, // name
			exchange.Type, // type
			true,          // durable
			false,         // autoDelete
			false,         // internal
			false,         // noWait
			nil,           // arguments
		)
		if err != nil {
			return fmt.Errorf("on Session.Connect: %v", err)
		}

		// create queues if not exists
		for _, queue := range exchange.Queues {
			// declare queue
			if _, err := s.ch.QueueDeclare(
				queue.Name, // name
				true,       // durable
				false,      // autoDelete
				false,      // exclusive
				false,      // noWait
				nil,        // arguments
			); err != nil {
				return fmt.Errorf("on Session.Connect: %v", err)
			}

			// bind the queue to the exchange
			if err := s.ch.QueueBind(
				queue.Name,     // name
				queue.RouteKey, // key
				exchange.Name,  // exchange
				false,          // noWait
				nil,            // arguments
			); err != nil {
				return fmt.Errorf("on Session.Connect: %v", err)
			}
		}
	}

	// reconnect if the connection closed
	go func() {
		// when the connection is closed
		if err, ok := <-conn.NotifyClose(make(chan *amqp.Error)); ok {
			console.Error("connection(%q) closed: %v", s.config.Addr, err.Error())
		} else {
			return // gracefule close
		}

		// try reconnect
		err := s.reconnect()
		if err != nil {
			console.Fatal("%v", err)
		}

		// notify that the connection is reconstructed
		s.connNotifier.Source <- struct{}{}
	}()

	return nil
}

func (s *Session) reconnect() error {
	// try reconnect 30 times
	for trial := 1; trial <= 30; trial++ {
		console.Error("try reconnect %d times...", trial)

		err := s.Connect()
		if err == nil {
			console.Info("connection(%q) recovered", s.config.Addr)
			return nil
		}

		console.Error("failed to reconnect: %v", err)
		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("failed to reconnect(%q)", s.config.Addr)
}

// Consume registers handler for the queue with the exchange. The handler must
// return bool if the handling success, as either Ack or Nack will be sent to
// RabbitMQ depending on the bool result.
func (s *Session) Consume(
	exchange, queue string,
	handler func([]byte) bool,
) error {
	// find target exchange and queue
	exCfg, ok := s.config.Exchanges[exchange]
	if !ok {
		return fmt.Errorf("undefined exchange(%q)", exchange)
	}
	qCfg, ok := exCfg.Queues[queue]
	if !ok {
		return fmt.Errorf("undefined queue(%q)", queue)
	}

	// open delivery channel from queue
	s.mu.Lock()
	delivery, err := s.ch.Consume(
		qCfg.Name, // queue
		"",        // consumer
		false,     // autoAck
		false,     // exclusive
		false,     // noLocal
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		s.mu.Unlock()
		return fmt.Errorf("on Consume: %v", err)
	}
	s.mu.Unlock()

	// notifier from reconnection event
	reconn := s.connNotifier.AddSubscriber()

	go func() {
		for {
			select {
			case <-reconn:
				// restart Consume after reconnect
				s.connNotifier.RemoveSubscriber(reconn)
				if err := s.Consume(exchange, queue, handler); err != nil {
					console.Fatal("failed in Consume: %v", err)
				}

			case d := <-delivery:
				// handle message from RabbitMQ
				if handler(d.Body) {
					if err := d.Ack(false); err != nil {
						console.Error("failed Ack: %v", err)
					}
				} else {
					if err := d.Nack(false, true); err != nil {
						console.Error("failed Nack: %v", err)
					}
				}
			}
		}
	}()

	return nil
}

// Publish publishes message to exchange with routing key.
func (s *Session) Publish(exchange, key string, msg []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.ch.Publish(
		exchange, // exchange
		key,      // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         msg,
			DeliveryMode: amqp.Transient,
		},
	); err != nil {
		return fmt.Errorf("on Session.Publish: %v", err)
	}

	return nil
}
