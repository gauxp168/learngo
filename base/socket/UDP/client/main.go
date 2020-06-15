package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp",nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("dail udp failed, error:", err)
		return
	}
	defer conn.Close()
	sendData := []byte("hello server")
	_, err = conn.Write(sendData)
	if err != nil {
		fmt.Println("sned data udp failed, error:", err)
		return
	}
	data := make([]byte, 4096)
	n, addr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read udp failed, error:", err)
		return
	}
	fmt.Printf("RECV :%v  addr:%v   count:%v\n", string(data[:n]), addr, n)
}






























