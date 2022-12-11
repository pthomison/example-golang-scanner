package hack

import (
	"fmt"
	"net"
	"time"

	"github.com/c-robinson/iplib"
	"github.com/jreisinger/nmapservices"
	"github.com/pthomison/errcheck"
)

var (
	TCPTimeout = 2 * time.Second
)

func Elapsed(start time.Time) {
	fmt.Printf("Time Elapsed: %v\n", time.Now().Sub(start))
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

func GetPopularPorts(limit int) []int {
	ports := make([]int, limit)

	services, err := nmapservices.Get()
	errcheck.Check(err)
	topNmapPorts := services.Tcp().Top(limit)

	for i, v := range topNmapPorts {
		ports[i] = int(v.Port)
	}

	return ports
}

func RangeNetwork(network net.IPNet) []net.IP {
	ips := make([]net.IP, 0)

	ip := network.IP

	for {
		if network.Contains(ip) {
			ips = append(ips, ip)
			ip = iplib.IncrementIP4By(ip, uint32(1))
		} else {
			break
		}
	}

	return ips
}
