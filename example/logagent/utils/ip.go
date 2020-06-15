package utils

import (
	"net"
	"strings"
)

func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	addr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(addr.IP.String(), ":")[0]
	return
}
