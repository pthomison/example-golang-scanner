package main

import (
	"fmt"
	"time"
)

const (
	WorkerCount     = 100
	TCPConnPoolSize = 100

	PortRangeStart = 1
	PortRangeEnd   = 100

	TCPTimeout = 100 * time.Millisecond

	TCPPortsCount = 500
)

var (
// Ports = services.Tcp().Top(500)
)

func main() {

	privateNetworks := findPrivateIPV4Networks()
	fmt.Printf("%+v\n", scanNetwork(privateNetworks[0]))

	// for i, s := range services.Tcp().Top(500) {
	// 	fmt.Printf("%2d. %5d # %s\n", i+1, s.Port, s.Name)
	// }

}
