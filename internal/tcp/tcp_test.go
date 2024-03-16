package tcp

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPacketHash(t *testing.T) {
	packet1 := TCPPackt{
		SrcIP:   net.ParseIP("192.168.1.1"),
		DstIP:   net.ParseIP("10.0.0.1"),
		SrcPort: 12345,
		DstPort: 80,
	}

	packet2 := TCPPackt{
		SrcIP:   net.ParseIP("10.0.0.1"),
		DstIP:   net.ParseIP("192.168.1.1"),
		SrcPort: 80,
		DstPort: 12345,
	}

	// Check that the hash values are equal
	assert.Equal(t, packet1.PacketHash(), packet2.PacketHash(), "Hash values should be equal for forward and backward packets")

}
