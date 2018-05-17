package main

import (
	"sync"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

// Backplane is a go struct pretending to be a pubsub service
type Backplane struct {
	sync.Mutex
	// Map of topic -> map of subscriber ID -> channel
	topics map[string]map[string]chan interface{}
}

func NewBackplane() *Backplane {
	return &Backplane{
		topics: make(map[string]map[string]chan interface{}),
	}
}

func (b *Backplane) Publish(topic string, message interface{}) {
	b.Lock()
	defer b.Unlock()
	listeners, ok := b.topics[topic]
	if ok {
		// broadcast
		for _, listener := range listeners {
			listener <- message
		}
	}
}

func (b *Backplane) Subscribe(ctx context.Context, topic string, callback func(message interface{})) {
	subID := uuid.NewV4().String()
	listener := make(chan interface{})

	listeners, ok := b.topics[topic]
	if !ok {
		listeners = make(map[string]chan interface{})
		b.topics[topic] = listeners
	}
	listeners[subID] = listener

	// Listener thread
	go func() {
		// Cleanup on termination
		defer func() {
			b.Lock()
			defer b.Unlock()
			delete(listeners, subID)
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-listener:
				callback(message)
			}
		}
	}()
}
