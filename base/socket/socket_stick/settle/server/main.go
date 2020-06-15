package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"simqo.com/mygospace/learngo/base/socket/socket_stick/settle/proto"
)

func process(conn net.Conn)  {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for  {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, error:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:6789")
	if err != nil {
		fmt.Println("tcp listen failed, error:", err)
		return
	}
	defer listener.Close()
	for  {
		conn, err := listener.Accept()
		if err != nil  {
			fmt.Println("tcp accept failed, error:", err)
			return
		}
		go process(conn)
	}
}
