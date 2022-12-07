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

	utils.RangeNetwork(&IPNetwork, func(ip net.IP) {
		scanIP(ip)
	})

	// fmt.Printf("Scanning IP (%v): ", IPTarget)

	// for _, port := range Ports {
	// 	go func(port int) {
	// 		wg.Add(1)
	// 		if utils.IsPortOpen(IPTarget, port) {
	// 			fmt.Printf("%v ", port)
	// 		}
	// 		wg.Done()
	// 	}(port)
	// }

	// wg.Wait()

	// fmt.Println()
}

func scanIP(ip net.IP) {
	fmt.Printf("Scanning IP (%v): ", ip)

	Ports := utils.GetPopularPorts(PortsCount)

	for _, port := range Ports {
		go func(port int) {
			wg.Add(1)
			if utils.IsPortOpen(ip, port) {
				fmt.Printf("%v ", port)
			}
			wg.Done()
		}(port)
	}

	wg.Wait()

	fmt.Println()
}
