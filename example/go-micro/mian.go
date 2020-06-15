package main

import (
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	"net/http"
)

func main()  {
	etcd.NewRegistry()
	server := web.NewService(web.Address(":8080"))
	server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world!"))
	})
	server.Run()
}
