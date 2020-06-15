package main

import (
	"fmt"
	"net"
	"simqo.com/mygospace/learngo/base/socket/socket_stick/settle/proto"
)

func main() {
	conn, err := net.Dial("tcp", "1297.0.0.1:6789")
	if err != nil {
		fmt.Println("tcp dail failed, error:", err)
		return
	}
	defer conn.Close()
	for i:= 0; i< 22; i++ {
		//data := []byte("你好 ！ hello . how are you?")
		msg := "你好 ！ hello . how are you?"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("client send data failed, error:", err)
			continue
		}
	}
}
