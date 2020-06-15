package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width ,Height int
} 

func main() {
	client, err := jsonrpc.Dial("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	p := Params{50, 100}
	ret := 0
	err = client.Call("Rect.Area", p, &ret)
	if err != nil {
		log.Println(err)
	}
	err = client.Call("Rect.Perimeter", p, &ret)
	if err != nil {
		log.Println(err)
	}
}
