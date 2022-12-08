package utils

import (
	"fmt"
	"net"
	"sync"
)

func ScanIP(ip net.IP, ports []int) string {
	wg := new(sync.WaitGroup)

	output := fmt.Sprintf("Scanning IP (%v):", ip)

	// Ports := GetPopularPorts(PortsCount)

	for _, port := range ports {
		go func(port int) {
			wg.Add(1)
			if IsPortOpen(ip, port) {
				output = fmt.Sprintf("%v %v", output, port)
			}
			wg.Done()
		}(port)
	}

	wg.Wait()

	return output
}
