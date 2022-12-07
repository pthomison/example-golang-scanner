package main

import (
	"fmt"
	"net"

	"github.com/pthomison/example-golang-scanner/utils"
)

var (
	PortsCount = 10000
	Ports      = utils.GetPopularPorts(PortsCount)

	IPTarget = net.IPv4(192, 168, 0, 1)
)

func main() {

	fmt.Printf("Scanning IP (%v): ", IPTarget)

	for _, port := range Ports {
		if utils.IsPortOpen(IPTarget, port) {
			fmt.Printf("%v ", port)
		}
	}

	fmt.Println()

}
