package server_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/perryd01/fh6t/internal/server"
	"github.com/perryd01/fh6t/internal/telemetry"
)

func TestSSEHub_DeliversPacketToSubscriber(t *testing.T) {
	ch := make(chan telemetry.Packet, 1)
	hub := server.NewSSEHub(ch)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go hub.Run(ctx)

	client := hub.Subscribe()
	defer hub.Unsubscribe(client)

	ch <- telemetry.Packet{IsRaceOn: 1, TimestampMS: 99}

	select {
	case msg := <-client:
		var p telemetry.Packet
		if err := json.Unmarshal(msg, &p); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if p.TimestampMS != 99 {
			t.Errorf("TimestampMS: got %d, want 99", p.TimestampMS)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timed out waiting for packet")
	}
}

func TestSSEHub_UnsubscribedClientReceivesNothing(t *testing.T) {
	ch := make(chan telemetry.Packet, 1)
	hub := server.NewSSEHub(ch)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go hub.Run(ctx)

	client := hub.Subscribe()
	hub.Unsubscribe(client)

	ch <- telemetry.Packet{IsRaceOn: 1}

	select {
	case <-client:
		t.Fatal("unsubscribed client should not receive packets")
	case <-time.After(50 * time.Millisecond):
		// correct: nothing received
	}
}
