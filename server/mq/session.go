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

// TryConnect tries Connect try times with given interval. Returns error if
// still not connected even after trials.
func (s *Session) TryConnect(try int, interval time.Duration) error {
	for trial := 0; trial < try; trial++ {
		if trial > 0 {
			console.Error("try reconnect %d times...", trial)
		}

		err := s.connect()
		if err == nil {
			return nil
		}
		console.Error("failed to TryConnect: %v", err)

		time.Sleep(interval)
	}

	return fmt.Errorf("failed to connect(%q)", s.config.Addr)
}

// connect establishes connection, channel, exchange, queue of RabbitMQ.
func (s *Session) connect() error {
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
			declared, err := s.ch.QueueDeclare(
				queue.Name, // name
				true,       // durable
				false,      // autoDelete
				false,      // exclusive
				false,      // noWait
				nil,        // arguments
			)
			if err != nil {
				return fmt.Errorf("on Session.Connect: %v", err)
			}

			// bind the queue to the exchange
			if err := s.ch.QueueBind(
				declared.Name,  // name
				queue.RouteKey, // key
				exchange.Name,  // exchange
				false,          // noWait
				nil,            // arguments
			); err != nil {
				return fmt.Errorf("on Session.Connect: %v", err)
			}

			// save queue name created by RabbitMQ
			queue.realName = declared.Name
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
		err := s.TryConnect(40, 3000*time.Millisecond)
		if err != nil {
			console.Fatal("on reconver connection: %v", err)
		}
		console.Info("connection(%q) recovered", s.config.Addr)

		// notify that the connection is reconstructed
		s.connNotifier.Source <- struct{}{}
	}()

	return nil
}

// Consume registers handler for the queue with the exchange. The handler must
// return true if the handling is successful. If it was unsuccessful, return
// false with error.
func (s *Session) Consume(
	exchange, queue string,
	handler func(*amqp.Delivery) (bool, error),
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
		qCfg.realName, // queue
		"",            // consumer
		false,         // autoAck
		false,         // exclusive
		false,         // noLocal
		false,         // noWait
		nil,           // arguments
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
				if ok, err := handler(&d); !ok {
					console.Warning("failed to handle delivery: %v", err)
				}
				d.Ack(false)
			}
		}
	}()

	return nil
}

// Publish publishes message to exchange with routing key.
func (s *Session) Publish(exchange, key, sender string, msg []byte) error {
	var header amqp.Table
	if sender != "" {
		header = amqp.Table{
			server.Sender: sender,
		}
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.ch.Publish(
		exchange, // exchange
		key,      // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			Headers:      header,
			ContentType:  "application/json",
			Body:         msg,
			DeliveryMode: amqp.Transient,
		},
	); err != nil {
		return fmt.Errorf("on Session.Publish: %v", err)
	}

	return nil
}
