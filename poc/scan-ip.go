package main

import (
	"fmt"
	"time"

	"github.com/pthomison/example-golang-scanner/poc/hack"
)

var (
	PortRangeStart = 1
	PortRangeEnd   = 10000
)

func main() {

	start := time.Now()

	fmt.Printf("Scanning IP (%v): ", hack.IPAddr)

	for port := PortRangeStart; port <= PortRangeEnd; port++ {

		if hack.IsPortOpen(hack.IPAddr, port) {
			fmt.Printf("%v ", port)
		}
	}

	fmt.Println()

	hack.Elapsed(start)

}
