package server

import "sync"

// Broadcaster broadcast a single event to multiple channels.
type Broadcaster struct {
	// Source is broadcasted to subscribers. Close Source to stop broadcaster.
	Source chan interface{}

	mu          sync.Mutex
	subscribers map[chan interface{}]struct{}
}

// NewBroadcaster creates Broadcaster.
func NewBroadcaster() *Broadcaster {
	br := &Broadcaster{
		Source:      make(chan interface{}),
		subscribers: make(map[chan interface{}]struct{}),
	}
	go br.run()
	return br
}

// AddSubscriber adds subscriber and return broadcasting channel.
func (b *Broadcaster) AddSubscriber() chan interface{} {
	ch := make(chan interface{})

	b.mu.Lock()
	defer b.mu.Unlock()

	b.subscribers[ch] = struct{}{}
	return ch
}

// RemoveSubscriber close ch and remove subscriber.
func (b *Broadcaster) RemoveSubscriber(ch chan interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.subscribers[ch]; ok {
		close(ch)
		delete(b.subscribers, ch)
	}
}

func (b *Broadcaster) run() {
	for val := range b.Source {
		b.mu.Lock()
		for ch := range b.subscribers {
			ch <- val
		}
		b.mu.Unlock()
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	for ch := range b.subscribers {
		close(ch)
		delete(b.subscribers, ch)
	}
	b.subscribers = nil
}
