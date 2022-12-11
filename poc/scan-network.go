package main

import (
	"fmt"
	"time"

	"github.com/pthomison/example-golang-scanner/poc/hack"
)

func main() {

	start := time.Now()

	ips := hack.RangeNetwork(hack.IPNetwork)

	for _, ip := range ips {
		results := hack.ScanIP(ip, hack.Ports)

		if len(results.OpenPorts) > 0 {
			fmt.Println(results)
		}

	}

	hack.Elapsed(start)

}
