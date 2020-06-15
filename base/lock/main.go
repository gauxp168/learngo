package main

import (
	"fmt"
	"sync"
	"time"
)

// 并发安全和锁
// 在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态

// 1. 互斥锁
// 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。
// Go语言中使用sync包的Mutex类型来实现互斥锁
var x int
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock  sync.RWMutex

func add()  {
	for i := 0; i<100; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

// 2. 读写互斥锁
// 互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。
// 读写锁在Go语言中使用sync包中的RWMutex类型。
//
//读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

func write()  {
	rwlock.Lock()
	x += 1
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock()
	wg.Done()
}

func read()  {
	rwlock.RLock()
	//rwlock.Lock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	//rwlock.Unlock()
	wg.Done()
}

func main() {
	//wg.Add(2)
	//go add()
	//go add()
	//wg.Wait()
	//fmt.Println(x)

	start := time.Now()
	for i:=0; i <100; i++ {
		wg.Add(1)
		go write()
	}
	for i:=0;i<100 ;i++  {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
	// 1.0770284s
	// 1.2802867s
}
