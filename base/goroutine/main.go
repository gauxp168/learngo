package main

import (
	"fmt"
	"runtime"
	"sync"
)

// goroutine
//goroutine 奉行通过通信来共享内存，而不是共享内存来通信
//Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。
//
//一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。


var wg sync.WaitGroup

func hello(i int)  {
	wg.Done()
	fmt.Println("hello", i)
}

func main() {

	runtime.GOMAXPROCS(1) // // 默认CPU的逻辑核心数，默认跑满整个CPU
	fmt.Println(runtime.NumCPU())

	for i := 0; i <200; i++ {
		wg.Add(1)
		//go hello(i)
		go func(i int) {
			wg.Done()
			fmt.Println(i)
		}(i)
	}
	fmt.Println("main")
	//time.Sleep(time.Second)
	wg.Wait()
}
