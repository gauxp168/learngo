package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	defer client.Close()
	// put 操作
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = client.Put(ctx, "key1", "value1")
	if err != nil {
		fmt.Printf("put to etcd failed, error:%v\n", err)
		return
	}
	cancel()

	// get 操作
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	response, err := client.Get(ctx, "key1")
	if err != nil {
		fmt.Printf("get etcd data failed, error:%v\n", err)
		return
	}
	for _, value := range response.Kvs {
		fmt.Printf("%s:%s\n", value.Key, value.Value)
	}

}
