package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/pthomison/example-golang-scanner/utils"
)

var (
	PortsCount = 1000

	IPNetwork = net.IPNet{
		IP:   net.IPv4(192, 168, 1, 0),
		Mask: []byte{255, 255, 255, 0},
	}

	wg = new(sync.WaitGroup)
)

func main() {

	ips := utils.RangeNetwork(IPNetwork)
	ports := utils.GetPopularPorts(PortsCount)

	for _, ip := range ips {
		results := utils.ScanIP(ip, ports)

		if len(results.OpenPorts) > 0 {
			fmt.Println(results)
		}

	}
}
