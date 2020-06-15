package main

import (
	"google.golang.org/grpc"
	"net/http"
	"simqo.com/mygospace/learngo/example/grpc/base/services"
)

func main()  {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	/*listen, _ := net.Listen("tcp", ":8080")
	rpcServer.Serve(listen)*/

	// http 服务模式
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		rpcServer.ServeHTTP(writer,request)
	})
	httpServer := &http.Server{
		Addr:":8080",
		Handler:mux,
	}
	httpServer.ListenAndServeTLS("","")
}
