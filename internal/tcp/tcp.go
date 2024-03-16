package tcp

import (
	"encoding/binary"
	"net"

	"github.com/mohamdafzal06/tcplat/internal/fnv"
)

// Packet contains the attributes of the packet that we desire.
type Packet struct {
	//SrcIP     netaddr.IP
	SrcIP   net.IP
	SrcPort uint16
	//DestIP    netaddr.IP
	DestIP    net.IP
	DestPort  uint16
	Syn       bool
	Ack       bool
	TimeStamp uint64
}

// TCPPackt is the attributes of a packet that we desired.
type TCPPackt struct {
	SrcIP   net.IP
	DstIP   net.IP
	SrcPort uint16
	DstPort uint16
	Syn     bool
	Ack     bool
}

// PacketHash, hash a packet for checksum
func (pkt TCPPackt) PacketHash() uint64 {
	tmp := make([]byte, 2)
	var src, dest []byte

	binary.BigEndian.PutUint16(tmp, pkt.SrcPort)
	srcAddr := pkt.SrcIP.To16()
	src = append(srcAddr[:], tmp...)

	binary.BigEndian.PutUint16(tmp, pkt.DstPort)
	destAddr := pkt.DstIP.To16()
	dest = append(destAddr[:], tmp...)

	return (fnv.Hash(src) + fnv.Hash(dest)) * fnv.Prime
}
