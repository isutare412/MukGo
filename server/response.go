package server

import (
	"fmt"
	"math/rand"
	"sync"
)

// ChannelMap is a multiplexer for response handlers from RabbitMQ RPC.
type ChannelMap struct {
	// IDGet generates unique correlation id with [32]byte.
	IDGet <-chan string

	handlers map[string]chan Packet
	mu       sync.Mutex
}

// NewChannelMap creates ResponseMux safely.
func NewChannelMap() *ChannelMap {
	mux := &ChannelMap{
		handlers: make(map[string]chan Packet),
	}

	// run id generator
	var idCh = make(chan string)
	go func(ch chan<- string) {
		var bytes [32]byte
		for {
			for i := range bytes {
				bytes[i] = byte(65 + rand.Intn(26)) // A-Z
			}
			ch <- string(bytes[:])
		}
	}(idCh)
	mux.IDGet = idCh

	return mux
}

// Register handler with given id.
func (m *ChannelMap) Register(
	id string, ch chan Packet,
) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.handlers[id]; ok {
		return fmt.Errorf("handler with id(%s) already exists", id)
	}
	m.handlers[id] = ch

	return nil
}

// Pop get and remove handlers with given id.
func (m *ChannelMap) Pop(id string) chan Packet {
	m.mu.Lock()
	defer m.mu.Unlock()

	h, ok := m.handlers[id]
	if !ok {
		return nil
	}
	delete(m.handlers, id)
	return h
}

// Get retrieves handler with given id.
func (m *ChannelMap) Get(id string) chan Packet {
	m.mu.Lock()
	defer m.mu.Unlock()

	h, ok := m.handlers[id]
	if !ok {
		return nil
	}
	return h
}
