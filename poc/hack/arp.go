package hack

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/j-keck/arping"
	"github.com/pthomison/errcheck"
)

type ARPResult struct {
	IP           net.IP
	HWAddr       net.HardwareAddr
	ResponseTime time.Duration
	Online       bool
}

func ARPScan(network net.IPNet) []*ARPResult {

	c := make(chan *ARPResult, len(RangeNetwork(network)))
	ouput := make([]*ARPResult, 0)

	NetworkARP(network, func(result *ARPResult) {
		c <- result
	})

	for len(c) > 0 {
		ouput = append(ouput, <-c)
	}

	return ouput

}

func NetworkARP(network net.IPNet, callback func(*ARPResult)) {
	ips := RangeNetwork(network)
	wg := new(sync.WaitGroup)

	for _, ip := range ips {
		wg.Add(1)
		go func(addr net.IP) {
			result := sendARP(addr, 0)
			if result.Online {
				callback(result)
			}
			wg.Done()
		}(ip)

		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
}

func sendARP(ip net.IP, retryCount int) *ARPResult {
	result := &ARPResult{
		IP:           ip,
		HWAddr:       nil,
		ResponseTime: 0,
		Online:       false,
	}

	hwAddr, dur, err := arping.Ping(ip)
	if err == arping.ErrTimeout {
		return result
	} else if err != nil {
		// fmt.Printf("Retrying %v: error(%v)\n", ip, err)
		retryCount += 1
		for i := 0; i < retryCount; i++ {
			time.Sleep(1000 * time.Millisecond)
		}

		if retryCount > 5 {
			errcheck.Check(errors.New(fmt.Sprintf("Unable to ARP %v", ip)))
		}

		return sendARP(ip, retryCount)
	}

	result.Online = true
	result.ResponseTime = dur
	result.HWAddr = hwAddr

	return result
}
