package utils

import (
	"fmt"
	"net"
	"time"
)

var (
	TCPTimeout = 100 * time.Millisecond
)

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
