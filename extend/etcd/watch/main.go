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
		fmt.Printf("connect etcd failed, error:%v\n", err)
		return
	}
	defer  client.Close()
	watch := client.Watch(context.Background(), "key1")
	for value := range watch {
		for _, ev := range value.Events {
			fmt.Printf("type:%v  key:%v  value:%v \n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
