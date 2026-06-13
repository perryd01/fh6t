package session_test

import (
	"testing"
	"time"

	"github.com/perryd01/fh6t/internal/session"
	"github.com/perryd01/fh6t/internal/telemetry"
)

func TestMemoryStore_SaveAndGet(t *testing.T) {
	s := session.Session{
		ID:        "lap-1",
		StartedAt: time.Now(),
		Packets:   []telemetry.Packet{{IsRaceOn: 1}},
	}

	ms := session.NewMemoryStore()
	err := ms.Save(s)
	if err != nil {
		t.Fatal("unexpected error on session save")
	}

	getSession, err := ms.Get(s.ID)
	if err != nil {
		t.Fatal("unexpected error on session get")
	}

	if getSession.ID != "lap-1" {
		t.Fatal("ID not matching the session")
	}

}

func TestMemoryStore_List(t *testing.T) {
	memoryStore := session.NewMemoryStore()
	sessions, err := memoryStore.List()
	if err != nil {
		t.Fatal("unexpected error")
	}
	if len(sessions) != 0 {
		t.Fatal("sessions lenght is not null")
	}
	newSession := session.Session{
		ID:        "lap-1",
		StartedAt: time.Now(),
		Packets:   []telemetry.Packet{},
	}

	memoryStore.Save(newSession)

	sessions, err = memoryStore.List()
	if err != nil {
		t.Fatal("unexpected error")
	}
	if len(sessions) != 1 {
		t.Fatal("sessions length should be 1 after one save")
	}

}

func TestMemoryStore_GetNotFound(t *testing.T) {
	memoryStore := session.NewMemoryStore()

	_, err := memoryStore.Get("not-exists")
	if err == nil {
		t.Fatal("get not found should return an error")
	}

}
