package main

import (
	"fmt"
	"net"
	"time"
)

const (
	WorkerCount     = 100
	TCPConnPoolSize = 100

	// PortRangeStart = 1
	// PortRangeEnd   = 100

	TCPTimeout = 100 * time.Millisecond

	TCPPortsCount = 500
)

func main() {

	privateNetworks := findPrivateIPV4Networks()
	fmt.Printf("%+v\n", scanNetwork(privateNetworks[0]))

	fmt.Printf("%+v\n", ScanIP(net.IPv4(100, 66, 31, 77)))

	// ip, ipnet, err := ParseCIDR(s string)

}
