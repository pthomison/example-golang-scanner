package main

import (
	"fmt"
	"net"
	"regexp"

	"github.com/c-robinson/iplib"
	"github.com/pthomison/errcheck"
)

func isIPV4(ip net.IP) bool {
	ipv4_regexp := regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)
	return ipv4_regexp.MatchString(ip.String())
}

func findPrivateIPV4Networks() []*net.IPNet {
	privateNetworks := make([]*net.IPNet, 0)

	interfaces, err := net.Interfaces()
	errcheck.Check(err)

	for _, iface := range interfaces {
		ifaceAddrs, err := iface.Addrs()
		errcheck.Check(err)

		for _, addr := range ifaceAddrs {
			ip, ipnet, err := net.ParseCIDR(addr.String())
			errcheck.Check(err)

			if ip.IsPrivate() && isIPV4(ip) {
				privateNetworks = append(privateNetworks, ipnet)
			}
		}

	}

	return privateNetworks

}

func IsPortOpen(ip net.IP, port int) bool {
	addr := fmt.Sprintf("%v:%v", ip.String(), port)
	dialer := &net.Dialer{Timeout: TCPTimeout}
	conn, err := dialer.Dial("tcp4", addr)
	if err != nil {
		return false
	} else {
		conn.Close()
		return true
	}
}

func rangeNetwork(network *net.IPNet, callback func(net.IP)) {
	ip := iplib.IncrementIP4By(network.IP, uint32(1))

	for {
		if network.Contains(ip) {
			callback(ip)

			ip = iplib.IncrementIP4By(ip, uint32(1))
		} else {
			break
		}
	}
}
