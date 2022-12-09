package common

import (
	"net"

	"github.com/pthomison/example-golang-scanner/utils"
)

var (
	PortsCount = 10000

	IPAddr  = net.IPv4(192, 168, 1, 1)
	NetMask = []byte{255, 255, 255, 0}

	Ports = utils.GetPopularPorts(PortsCount)

	IPNetwork = net.IPNet{
		IP:   IPAddr,
		Mask: NetMask,
	}
)
