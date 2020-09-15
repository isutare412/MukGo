package mq

import (
	"fmt"
	"sync"
	"time"

	"github.com/isutare412/MukGo/server/console"
	"github.com/streadway/amqp"
)

// NewSession establishes new Session between server and message queue.
func NewSession(id, pw, addr string) (*Session, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s", id, pw, addr)

	session := new(Session)
	err := session.connect(url)
	if err != nil {
		return nil, fmt.Errorf("on NewSession: %v", err)
	}

	return session, nil
}

// Session keeps track of construected session between server and message queue.
// It is thread-safe.
type Session struct {
	*amqp.Channel
	conn *amqp.Connection
	mu   sync.Mutex
}

func (s *Session) connect(url string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// establish new connection
	var conn *amqp.Connection
	conn, err := amqp.Dial(url)
	if err != nil {
		return fmt.Errorf("on Session.connect: %v", err)
	}
	s.conn = conn

	// open new channel from connection
	var ch *amqp.Channel
	ch, err = conn.Channel()
	if err != nil {
		return fmt.Errorf("on Session.connect: %v", err)
	}
	s.Channel = ch

	// reconnect if the connection closed
	go func(url string) {
		err, byError := <-conn.NotifyClose(make(chan *amqp.Error))
		if !byError {
			return // graceful close
		}
		console.Error("connection(%q) closed: %v", url, err.Error())

		for trial := 1; trial <= 30; trial++ {
			console.Error("try reconnect %d times...", trial)

			err := s.connect(url)
			if err == nil {
				console.Info("connection(%q) recovered", url)
				return
			}

			console.Error("failed to reconnect: %v", err)
			time.Sleep(2 * time.Second)
		}
	}(url)

	return nil
}
