package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/pthomison/example-golang-scanner/utils"
)

var (
	PortsCount = 1000
	Ports      = utils.GetPopularPorts(PortsCount)

	IPTarget = net.IPv4(192, 168, 1, 1)

	wg = new(sync.WaitGroup)
)

func main() {
	fmt.Printf("Scanning IP (%v): ", IPTarget)

	for _, port := range Ports {
		go func(port int) {
			wg.Add(1)
			if utils.IsPortOpen(IPTarget, port) {
				fmt.Printf("%v ", port)
			}
			wg.Done()
		}(port)
	}

	wg.Wait()

	fmt.Println()
}
