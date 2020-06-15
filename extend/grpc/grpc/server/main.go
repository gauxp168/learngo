package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"simqo.com/mygospace/learngo/extend/grpc/protobuf"
)

// 1.需要监听
// 2.需要实例化gRPC服务端
// 3.在gRPC商注册微服务
// 4.启动服务端

type UserInfoService struct {

}

var u = UserInfoService{}

func (s *UserInfoService) GetUserInfo(ctx context.Context, in *proto.UserRequest ) (resp *proto.UserResponse, err error) {
	name := in.Name
	if name == "zx" {
		resp = &proto.UserResponse{
			Id:1,
			Name:"zx",
			Age:23,
			Hobby:[]string{"sing", "run"},
		}
	}
	return
}

func main() {
	addr := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	s := grpc.NewServer()
	proto.RegisterUserInfoServiceServer(s, &u)
	s.Serve(listener)
}
