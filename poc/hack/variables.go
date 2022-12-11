package hack

import (
	"net"
)

var (
	PortsCount = 10000

	IPAddr  = net.IPv4(192, 168, 1, 1)
	NetMask = []byte{255, 255, 255, 0}

	Ports = GetPopularPorts(PortsCount)

	IPNetwork = net.IPNet{
		IP:   IPAddr,
		Mask: NetMask,
	}
)
