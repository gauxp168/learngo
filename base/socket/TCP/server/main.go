package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
一个TCP服务端可以同时连接很多个客户端。
因为Go语言中创建多个goroutine实现并发非常方便和高效，
所以我们可以每建立一次链接就创建一个goroutine去处理。

TCP服务端程序的处理流程：
    1.监听端口
    2.接收客户端请求建立链接
    3.创建goroutine处理链接。
*/

func process(conn net.Conn)  {
	defer conn.Close()	//关闭连接
	for  {
		reader := bufio.NewReader(conn)		//读取数据
		var buf  [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, error:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")	//监听端口
	if err != nil {
		fmt.Println("tcp listen failed,error:", err)
		return
	}
	for  {
		conn, err := listener.Accept()	//建立连接
		if err != nil {
			fmt.Println("tcp accept failed, error:",err)
			continue
		}
		go process(conn)
	}
}
