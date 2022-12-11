package main

import (
	"fmt"
	"time"

	"github.com/pthomison/example-golang-scanner/poc/hack"
)

func main() {

	start := time.Now()

	fmt.Printf("Scanning IP (%v): ", hack.IPAddr)

	for _, port := range hack.Ports {
		if hack.IsPortOpen(hack.IPAddr, port) {
			fmt.Printf("%v ", port)
		}
	}

	fmt.Println()

	hack.Elapsed(start)

}
