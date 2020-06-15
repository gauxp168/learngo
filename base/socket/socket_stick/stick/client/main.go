package main

import (
	"fmt"
	"net"
)

// 黏包

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("tcp dail failed, error:", err)
		return
	}
	defer conn.Close()
	for i:=0; i<20; i++ {
		dataStr := `heloo, hello, how are you ?`
		conn.Write([]byte(dataStr))
	}
}
