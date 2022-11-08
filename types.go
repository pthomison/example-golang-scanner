package main

import (
	"bytes"
	"fmt"
	"net"
)

type NetworkScan struct {
	Network *net.IPNet
	Hosts   []*HostScan
}

func (n *NetworkScan) String() string {

	s := "\n"

	s = fmt.Sprintf("%vNetwork Scan Results\n", s)
	s = fmt.Sprintf("%v--------------------\n", s)
	s = fmt.Sprintf("%vNetwork: %v\n", s, n.Network)
	s = fmt.Sprintf("%vActive Host Count: %v\n", s, len(n.Hosts))

	rangeNetwork(n.Network, func(ip net.IP) {
		host := n.RetrieveHostScan(ip)
		if host != nil {
			s = fmt.Sprintf("%v%v\n", s, host)
		}
	})

	return s
}

func (n *NetworkScan) RetrieveHostScan(ip net.IP) *HostScan {

	for _, host := range n.Hosts {

		if (bytes.Compare(ip, host.IP)) == 0 {
			return host
		}
	}

	return nil
}

type HostScan struct {
	IP          net.IP
	ActivePorts []int
}

func (h *HostScan) String() string {
	return fmt.Sprintf("Host: %v - (%v)", h.IP, h.ActivePorts)
}

// type PortScan struct {
// 	IP         net.IP
// 	PortNumber int
// }
