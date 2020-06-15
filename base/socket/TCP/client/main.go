package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
一个TCP客户端进行TCP通信的流程如下：
    1.建立与服务端的链接
    2.进行数据收发
    3.关闭链接
*/

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("tcp dial failed, error:", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for  {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("read string failed, error:", err)
			continue
		}
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_,err = conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("conn write data failed, error:",err)
			return
		}
		var buf [512]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, error:",err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}




















