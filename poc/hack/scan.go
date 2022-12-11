package hack

import (
	"net"
	"sync"
)

type IPScanResults struct {
	IP        net.IP
	OpenPorts []int
}

func ScanIP(ip net.IP, ports []int) *IPScanResults {
	wg := new(sync.WaitGroup)

	results := &IPScanResults{
		IP:        ip,
		OpenPorts: make([]int, 0),
	}

	// Very unlikely to have an IP will all requested ports open, but don't want the channel to deadlock
	// Are there better options?
	// Should collector channels always be as wide as your number of workers? Potential output data?
	// How to start consuming data before the goroutine are done?
	portChan := make(chan int, len(ports))

	for _, port := range ports {
		wg.Add(1)

		go func(port int) {
			if IsPortOpen(ip, port) {
				portChan <- port
			}
			wg.Done()
		}(port)
	}

	wg.Wait()

	for len(portChan) > 0 {
		results.OpenPorts = append(results.OpenPorts, <-portChan)
	}

	return results
}
