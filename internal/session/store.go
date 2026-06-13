package session

import (
	"fmt"
	"sync"
)

type Store interface {
	Save(s Session) error
	List() ([]Session, error)
	Get(id string) (Session, error)
}

type MemoryStore struct {
	mu       sync.RWMutex
	sessions map[string]Session
}

func (m *MemoryStore) Save(s Session) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.sessions[s.ID] = s
	return nil
}

func (m *MemoryStore) List() ([]Session, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]Session, 0, len(m.sessions))
	for _, s := range m.sessions {
		out = append(out, s)
	}
	return out, nil
}

func (m *MemoryStore) Get(id string) (Session, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.sessions[id]
	if !ok {
		return Session{}, fmt.Errorf("session %q not found", id)
	}
	return s, nil
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{sessions: make(map[string]Session)}
}
