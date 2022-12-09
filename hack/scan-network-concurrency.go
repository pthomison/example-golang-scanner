package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/pthomison/example-golang-scanner/utils"
)

var (
	PortsCount = 10000

	IPNetwork = net.IPNet{
		IP:   net.IPv4(192, 168, 1, 0),
		Mask: []byte{255, 255, 255, 0},
	}

	wg = new(sync.WaitGroup)
)

func main() {

	ports := utils.GetPopularPorts(PortsCount)
	ips := utils.RangeNetwork(IPNetwork)

	resultsChan := make(chan *utils.IPScanResults, len(ips))

	for _, ip := range ips {
		wg.Add(1)

		go func(ip net.IP) {
			resultsChan <- utils.ScanIP(ip, ports)
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

}
