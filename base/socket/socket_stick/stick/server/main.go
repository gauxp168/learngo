package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// 黏包

func process(conn net.Conn)  {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for  {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read conn failed, error:", err)
			break
		}
		fmt.Printf("client data: %v\n", string(buf[:n]))
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("tcp send data failed, error:", err)
			break
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("tcp listen failed, error:", err)
		return
	}
	defer listener.Close()
	for  {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("tcp accept failed, error:", err)
			return
		}
		go process(conn)
	}
}
