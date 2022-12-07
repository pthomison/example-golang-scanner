package main

import (
	"fmt"
	"net"

	"github.com/pthomison/example-golang-scanner/utils"
)

var (
	PortRangeStart = 1
	PortRangeEnd   = 65535

	IPTarget = net.IPv4(192, 168, 0, 1)
)

func main() {

	fmt.Printf("Scanning IP (%v): ", IPTarget)

	for port := PortRangeStart; port <= PortRangeEnd; port++ {
		if utils.IsPortOpen(IPTarget, port) {
			fmt.Printf("%v ", port)
		}
	}

	fmt.Println()

}
