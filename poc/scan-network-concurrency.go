package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/pthomison/example-golang-scanner/poc/hack"
)

func main() {

	start := time.Now()

	wg := new(sync.WaitGroup)

	ips := hack.RangeNetwork(IPNetwork)

	resultsChan := make(chan *hack.IPScanResults, len(ips))

	for _, ip := range ips {
		wg.Add(1)

		go func(ip net.IP) {
			resultsChan <- hack.ScanIP(ip, hack.Ports)
			wg.Done()
		}(ip)

	}

	wg.Wait()

	for len(resultsChan) > 0 {
		result := <-resultsChan

		if len(result.OpenPorts) > 0 {
			fmt.Println(result)
		}

	}

	hack.Elapsed(start)

}
