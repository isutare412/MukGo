package mq

import (
	"encoding/json"
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
		for _, qcfg := range exchange.Queues {
			// declare queue
			queue, err := s.ch.QueueDeclare(
				qcfg.Name,       // name
				qcfg.Durable,    // durable
				qcfg.AutoDelete, // autoDelete
				false,           // exclusive
				false,           // noWait
				nil,             // arguments
			)
			if err != nil {
				return fmt.Errorf("on Session.Connect: %v", err)
			}

			// bind the queue to the exchange
			if err := s.ch.QueueBind(
				queue.Name,    // name
				qcfg.RouteKey, // key
				exchange.Name, // exchange
				false,         // noWait
				nil,           // arguments
			); err != nil {
				return fmt.Errorf("on Session.Connect: %v", err)
			}

			// save queue name created by RabbitMQ
			qcfg.realName = queue.Name
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
	excfg, ok := s.config.Exchanges[exchange]
	if !ok {
		return fmt.Errorf("undefined exchange(%q)", exchange)
	}
	qcfg, ok := excfg.Queues[queue]
	if !ok {
		return fmt.Errorf("undefined queue(%q)", queue)
	}

	// use generated name for anonymous queue
	qname := qcfg.Name
	if qname == "" {
		qname = qcfg.realName
	}

	// open delivery channel from queue
	s.mu.Lock()
	delivery, err := s.ch.Consume(
		qname, // queue
		"",    // consumer
		false, // autoAck
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // arguments
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

// Publish sends message to exchange with routing key.
func (s *Session) Publish(
	exchange, key, sender string, packet server.Packet,
) error {
	pub, err := newPublishing(packet, sender, "", "")
	if err != nil {
		return fmt.Errorf("on Session.Publish: %v", err)
	}
	return s.publish(exchange, key, pub)
}

// Reply sends reply message RPC request.
func (s *Session) Reply(
	exchange, key, sender, correlationID string, packet server.Packet,
) error {
	pub, err := newPublishing(packet, sender, "", correlationID)
	if err != nil {
		return fmt.Errorf("on Session.Publish: %v", err)
	}
	return s.publish(exchange, key, pub)
}

// RPC publishes message with reply request. It is used for Remote Procedure
// Call.
func (s *Session) RPC(
	exchange, key, sender, replyQueue, correlationID string,
	packet server.Packet,
) error {
	pub, err := newPublishing(packet, sender, replyQueue, correlationID)
	if err != nil {
		return fmt.Errorf("on Session.RPC: %v", err)
	}
	return s.publish(exchange, key, pub)
}

func newPublishing(
	packet server.Packet,
	sender string,
	replyQueue string,
	corrID string,
) (*amqp.Publishing, error) {
	// build header
	var header = amqp.Table{
		server.MsgType: int32(packet.Type()),
	}
	if sender != "" {
		header[server.Sender] = sender
	}

	// serialize packet
	msg, err := json.Marshal(packet)
	if err != nil {
		return nil, fmt.Errorf("on newPublishing: %v", err)
	}

	// build publishing
	pub := &amqp.Publishing{
		Headers:       header,
		ContentType:   "application/json",
		Body:          msg,
		DeliveryMode:  amqp.Transient,
		ReplyTo:       replyQueue,
		CorrelationId: corrID,
	}

	return pub, nil
}

func (s *Session) publish(exchange, key string, p *amqp.Publishing) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.ch.Publish(
		exchange, // exchange
		key,      // routing key
		false,    // mandatory
		false,    // immediate
		*p,       // publishing
	); err != nil {
		return fmt.Errorf("on Session.publish: %v", err)
	}

	return nil
}
