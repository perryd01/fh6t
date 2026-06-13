package session

import (
	"context"
	"fmt"
	"time"

	"github.com/perryd01/fh6t/internal/telemetry"
)

type Manager struct {
	store Store
	ch    <-chan telemetry.Packet
}

func NewManager(store Store, ch <-chan telemetry.Packet) *Manager {
	return &Manager{
		store: store,
		ch:    ch,
	}
}

func (m *Manager) Run(ctx context.Context) {
	var current *Session

	closeSession := func() {
		if current == nil {
			return
		}
		now := time.Now()
		current.EndedAt = &now
		m.store.Save(*current)
		current = nil
	}

	for {
		select {
		case <-ctx.Done():
			closeSession()
			return
		case p, ok := <-m.ch:
			if !ok {
				closeSession()
				return
			}
			if p.IsRaceOn == 0 {
				closeSession()
			} else {
				if current == nil {
					now := time.Now()
					current = &Session{
						ID:        fmt.Sprintf("session-%d", now.UnixNano()),
						StartedAt: now,
						LapNumber: &p.LapNumber,
					}
				} else if p.LapNumber != *current.LapNumber {
					closeSession()
					now := time.Now()
					current = &Session{
						ID:        fmt.Sprintf("session-%d", now.UnixNano()),
						StartedAt: now,
						LapNumber: &p.LapNumber,
					}
				}
				current.Packets = append(current.Packets, p)
			}
		}
	}
}
