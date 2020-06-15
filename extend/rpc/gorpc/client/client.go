package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width int
	Height int
}

func main() {
	// 1. 链接远程rpc服务
	conn, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	p := Params{50, 100}
	//2. 调用方法
	ret := 0
	err = conn.Call("Rect.Area", p, &ret)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("面积：", ret)
	err = conn.Call("Rect.Perimeter", p, &ret)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("周长：", ret)
}
