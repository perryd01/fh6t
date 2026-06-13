package session_test

import (
	"context"
	"testing"
	"time"

	"github.com/perryd01/fh6t/internal/session"
	"github.com/perryd01/fh6t/internal/telemetry"
)

func TestManager_RecordsPackectsInSession(t *testing.T) {
	store := session.NewMemoryStore()
	c := make(chan telemetry.Packet)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	manager := session.NewManager(store, c)

	go manager.Run(ctx)

	c <- telemetry.Packet{
		IsRaceOn:  1,
		LapNumber: 1,
	}

	c <- telemetry.Packet{
		IsRaceOn:  1,
		LapNumber: 1,
	}

	c <- telemetry.Packet{
		IsRaceOn:  1,
		LapNumber: 2,
	}

	c <- telemetry.Packet{
		IsRaceOn:  1,
		LapNumber: 2,
	}

	time.Sleep(20 * time.Millisecond)

	sessions, err := store.List()
	if err != nil {
		t.Fatal("unexpected error on store list")
	}

	if len(sessions) != 1 {
		t.Fatal("should only have one session")
	}

	s := sessions[0]

	if len(s.Packets) != 2 {
		t.Fatal("should contain 2 packets")
	}

}

func TestManager_ClosesSessionOnRaceOff(t *testing.T) {
	store := session.NewMemoryStore()
	c := make(chan telemetry.Packet)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	manager := session.NewManager(store, c)

	go manager.Run(ctx)

	c <- telemetry.Packet{
		IsRaceOn:  1,
		LapNumber: 1,
	}

	c <- telemetry.Packet{
		IsRaceOn:  0,
		LapNumber: 1,
	}

	time.Sleep(20 * time.Millisecond)

	sessions, err := store.List()
	if err != nil {
		t.Fatal("unexpected error")
	}

	sessionLength := len(sessions)
	if sessionLength != 1 {
		t.Fatalf("expected session length of 1, got %d", sessionLength)
	}

	s := sessions[0]

	if s.EndedAt == nil {
		t.Fatal("endedAt should not be nil")
	}

}
