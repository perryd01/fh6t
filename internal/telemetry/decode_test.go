package telemetry_test

import (
	"testing"

	"github.com/perryd01/fh6t/internal/telemetry"
)

func TestDecode_WrongSize(t *testing.T) {
	_, err := telemetry.Decode(make([]byte, 100))
	if err == nil {
		t.Fatal("expected error for wrong-size packet, got nil")
	}
}

func TestDecode_ValidPacket(t *testing.T) {
	data := make([]byte, 324)
	p, err := telemetry.Decode(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.IsRaceOn != 0 {
		t.Errorf("IsRaceOn: got %d, want 0", p.IsRaceOn)
	}
}
