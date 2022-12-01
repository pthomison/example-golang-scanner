package main

import (
	"net"
	"sync"

	"github.com/c-robinson/iplib"
	"github.com/jreisinger/nmapservices"
	"github.com/pthomison/errcheck"
)

func scanNetwork(network *net.IPNet) *NetworkScan {
	var wg sync.WaitGroup

	result := &NetworkScan{
		Network: network,
		Hosts:   make([]*HostScan, 0),
	}

	chanSize := iplib.Net4{IPNet: *network}.Count() + 1
	inputChannel := make(chan net.IP, chanSize)
	outputChannel := make(chan *HostScan, chanSize)

	for i := 0; i < WorkerCount; i++ {
		wg.Add(1)
		go ScanIPWorker(inputChannel, outputChannel, &wg)
	}

	rangeNetwork(network, func(ip net.IP) { inputChannel <- ip })

	close(inputChannel)

	wg.Wait()

	for len(outputChannel) > 0 {
		result.Hosts = append(result.Hosts, <-outputChannel)
	}

	return result
}

func ScanIP(ip net.IP) *HostScan {
	var wg sync.WaitGroup

	result := &HostScan{
		IP:          ip,
		ActivePorts: make([]int, 0),
	}

	services, err := nmapservices.Get()
	errcheck.Check(err)
	ports := services.Tcp().Top(TCPPortsCount)

	chanSize := TCPPortsCount + 1
	inputChan := make(chan int, chanSize)
	outputChan := make(chan int, chanSize)

	for i := 0; i < TCPConnPoolSize; i++ {
		wg.Add(1)
		go ScanPoolWorker(ip, inputChan, outputChan, &wg)
	}

	for _, port := range ports {
		inputChan <- int(port.Port)
	}

	close(inputChan)

	wg.Wait()

	for len(outputChan) > 0 {
		result.ActivePorts = append(result.ActivePorts, <-outputChan)
	}

	return result
}
