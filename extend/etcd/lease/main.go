package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect etcd failed, error:%v\n", err)
		return
	}
	defer client.Close()

	// 创建一个5秒的租约
	response, err := client.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Put(context.TODO(), "key1", "value1", clientv3.WithLease(response.ID))
	if err != nil {
		log.Fatalln(err)
	}

}
