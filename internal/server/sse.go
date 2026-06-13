package server

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/perryd01/fh6t/internal/telemetry"
)

type SSEHub struct {
	ch      <-chan telemetry.Packet
	mu      sync.Mutex
	clients map[chan []byte]struct{}
}

func NewSSEHub(ch <-chan telemetry.Packet) *SSEHub {
	return &SSEHub{
		ch:      ch,
		clients: make(map[chan []byte]struct{}),
	}
}

func (hub *SSEHub) Subscribe() chan []byte {
	c := make(chan []byte, 16)
	hub.mu.Lock()
	hub.clients[c] = struct{}{}
	hub.mu.Unlock()
	return c
}

func (hub *SSEHub) Unsubscribe(c chan []byte) {
	hub.mu.Lock()
	delete(hub.clients, c)
	hub.mu.Unlock()
}

func (hub *SSEHub) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case p, ok := <-hub.ch:
			if !ok {
				return
			}

			data, err := json.Marshal(p)
			if err != nil {
				continue
			}
			hub.mu.Lock()
			for c := range hub.clients {
				select {
				case c <- data:
				default:
				}
			}
			hub.mu.Unlock()

		}
	}
}
