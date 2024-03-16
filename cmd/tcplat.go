package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vishvananda/netlink"
)

func main() {
	var ifaceStr string
	flag.StringVar(&ifaceStr, "iface", "eth0", "interface to attach the probe to")
	flag.Parse()
	iface, err := netlink.LinkByName(ifaceStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	qdiscs, err := netlink.QdiscList(iface)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// qDiscAttr := netlink.QdiscAttrs{
	// 	LinkIndex: iface.Attrs().Index,
	// 	Handle:    netlink.MakeHandle(0xffff, 0),
	// 	Parent:    netlink.HANDLE_MIN_INGRESS,
	// }
}

// added by mamreza :)
func newClsQdisc() {
	link, err := netlink.LinkByName("eth0") // Replace "eth0" with your network interface name
	if err != nil {
		fmt.Printf("Error fetching link: %v\n", err)
		return
	}

	// Define the attributes for the new clsact Qdisc
	clsactAttrs := netlink.QdiscAttrs{
		LinkIndex: link.Attrs().Index,
		Handle:    netlink.HANDLE_CLSACT,  // clsact uses a specific handle
		Parent:    netlink.HANDLE_INGRESS, // clsact can be attached to ingress
	}

	// Create the clsact Qdisc
	clsact := &netlink.Clsact{
		QdiscAttrs: clsactAttrs,
	}

	// Add the clsact Qdisc to the link
	if err := netlink.QdiscAdd(clsact); err != nil {
		fmt.Printf("Error adding clsact Qdisc: %v\n", err)
		return
	}
}
