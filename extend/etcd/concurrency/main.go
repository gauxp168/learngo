package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"log"
)

func main() {
	endpoints := []string{"127.0.0.1:2379"}
	client, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	//创建两个单独的会话用来演示锁竞争
	s1, err := concurrency.NewSession(client)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	m1 := concurrency.NewMutex(s1, "/my-lock/")
	s2, err := concurrency.NewSession(client)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	m2 := concurrency.NewMutex(s2, "/my-lock/")

	// 会话s1获取锁
	err = m1.Lock(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("acquired lock for s1")
	m2locked := make(chan struct{})
	go func() {
		defer close(m2locked)
		if err := m2.Lock(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	if err := m1.Unlock(context.TODO()); err != nil{
		log.Fatal(err)
	}
	fmt.Println("released lock for s1")
	<-m2locked
	fmt.Println("acquired lock for s2")
}
