package main

import (
	"fmt"
	"net"
)

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 接受服务器返回信息
	go func() {
		for  {
			buf:=make([]byte, 4096)
			cnt, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("收到的信息长度：%d，收到的信息内容：%v\n",cnt, buf)
		}
	}()

	// 等待给服务器输入信息
	for  {
		var input sting
		fmt.Scanln(&input)
		conn.Write([]byte(input))
	}
}