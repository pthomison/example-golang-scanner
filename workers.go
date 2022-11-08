package main

import (
	"net"
	"sync"
)

func ScanIPWorker(inputChan chan net.IP, outputChan chan *HostScan, wg *sync.WaitGroup) {

	for inputIP := range inputChan {
		hostResults := ScanIP(inputIP)
		if len(hostResults.ActivePorts) > 0 {
			outputChan <- hostResults
		}
	}

	wg.Done()
}

func ScanPoolWorker(ip net.IP, inputChan chan int, outputChan chan int, wg *sync.WaitGroup) {

	for port := range inputChan {
		if IsPortOpen(ip, port) {
			outputChan <- port
		}
	}

	wg.Done()
}
