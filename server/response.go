package server

import (
	"fmt"
	"math/rand"
	"sync"
)

// HandleMap is a multiplexer for response handlers from RabbitMQ RPC.
type HandleMap struct {
	// IDGet generates unique correlation id with [32]byte.
	IDGet <-chan string

	handlers map[string]func(bool, Packet)
	mu       sync.Mutex
}

// NewHandleMap creates ResponseMux safely.
func NewHandleMap() *HandleMap {
	mux := &HandleMap{
		handlers: make(map[string]func(bool, Packet)),
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
func (m *HandleMap) Register(
	id string, handler func(bool, Packet),
) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.handlers[id]; ok {
		return fmt.Errorf("handler with id(%s) already exists", id)
	}

	m.handlers[id] = handler

	return nil
}

// Pop get and remove handlers with given id.
func (m *HandleMap) Pop(id string) func(bool, Packet) {
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
func (m *HandleMap) Get(id string) func(bool, Packet) {
	m.mu.Lock()
	defer m.mu.Unlock()

	h, ok := m.handlers[id]
	if !ok {
		return nil
	}
	return h
}
