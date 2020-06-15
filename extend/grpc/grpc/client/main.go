package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"simqo.com/mygospace/learngo/extend/grpc/protobuf"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常: %s \n", err)
	}
	defer conn.Close()
	client := proto.NewUserInfoServiceClient(conn)
	req := new(proto.UserRequest)
	req.Name = "zs"
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("响应异常： %s\n", err)
	}

	fmt.Println(response)
}
