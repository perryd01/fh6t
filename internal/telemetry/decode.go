package telemetry

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const packetSize = 324

func Decode(data []byte) (Packet, error) {
	if len(data) != packetSize {
		return Packet{}, fmt.Errorf("telemetry: packet size %d, want %d", len(data), packetSize)
	}

	var packet Packet
	err := binary.Read(bytes.NewReader(data), binary.LittleEndian, &packet)
	if err != nil {
		return Packet{}, err
	}

	return packet, nil
}
