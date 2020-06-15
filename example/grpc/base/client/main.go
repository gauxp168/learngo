package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"simqo.com/mygospace/learngo/example/grpc/base/services"
)

func main()  {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	client := services.NewProdServiceClient(conn)
	response, err := client.GetProdStock(context.Background(),
		&services.ProdRequset{ProdId: 12})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.ProdStock)
}
