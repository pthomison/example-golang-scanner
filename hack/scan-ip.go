package main

import (
	"fmt"

	"github.com/pthomison/example-golang-scanner/hack/common"
	"github.com/pthomison/example-golang-scanner/utils"
)

var (
	PortRangeStart = 1
	PortRangeEnd   = 10000
)

func main() {

	fmt.Printf("Scanning IP (%v): ", common.IPAddr)

	for port := PortRangeStart; port <= PortRangeEnd; port++ {

		// start := time.Now()

		if utils.IsPortOpen(common.IPAddr, port) {
			fmt.Printf("%v ", port)
		}

		// fmt.Printf("%v - %v\n", port, time.Now().Sub(start))
	}

	fmt.Println()

}
