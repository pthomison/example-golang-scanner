package main

import (
	"fmt"
	"time"

	"github.com/pthomison/example-golang-scanner/poc/hack"
)

func main() {
	start := time.Now()
	arpResults := hack.ARPScan(hack.IPNetwork)

	for _, arpResult := range arpResults {

		ip := arpResult.IP
		results := hack.ScanIP(ip, hack.Ports)

		if len(results.OpenPorts) > 0 {
			fmt.Println(results)
		}

	}

	hack.Elapsed(start)
}
