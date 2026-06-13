package session

import (
	"time"

	"github.com/perryd01/fh6t/internal/telemetry"
)

type Session struct {
	ID        string
	LapNumber *uint16
	StartedAt time.Time
	EndedAt   *time.Time
	Packets   []telemetry.Packet
}
