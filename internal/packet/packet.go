package packet

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// TCPv4ACK constrcut a IPv4 TCP ACK packet, and returns the slice.
func TCPv4ACK() []byte {
	buf := gopacket.NewSerializeBuffer()
	eth := &layers.Ethernet{
		SrcMAC:       net.HardwareAddr([]byte{0, 1, 2, 3, 4, 5}),
		DstMAC:       net.HardwareAddr([]byte{5, 4, 3, 2, 1, 0}),
		EthernetType: layers.EthernetTypeIPv4,
	}
	ip := &layers.IPv4{
		SrcIP:    net.IP{1, 2, 3, 4},
		DstIP:    net.IP{5, 6, 7, 8},
		Protocol: layers.IPProtocolTCP,
	}

	tcp := &layers.TCP{
		SrcPort: 57675,
		DstPort: 80,
		SYN:     false,
		ACK:     true,
	}
	err := eth.SerializeTo(buf, gopacket.SerializeOptions{})
	if err != nil {
		panic("cannot serialize the Ethernet part to buffer.")
	}

	err = ip.SerializeTo(buf, gopacket.SerializeOptions{})
	if err != nil {
		panic("cannot serialize the IP part to buffer.")
	}

	err = tcp.SerializeTo(buf, gopacket.SerializeOptions{})
	if err != nil {
		panic("cannot serialize the TCP part to buffer.")
	}
	return buf.Bytes()
}

func TCPv4SYN() []byte    { return []byte{} }
func TCPv4SYNACK() []byte { return []byte{} }
