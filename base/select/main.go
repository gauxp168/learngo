package main

import (
	"fmt"
	"time"
)

// select
// 1. select多路复用
// Go内置了select关键字，可以同时响应多个通道的操作。
//select的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。具体格式如下：
/*
select {
case <-chan1:
   // 如果chan1成功读到数据，则进行该case处理语句
case chan2 <- 1:
   // 如果成功向chan2写入数据，则进行该case处理语句
default:
   // 如果上面都没有成功，则进入default处理流程
}
select特点：
	1. select可以同时监听一个或多个channel，直到其中一个channel ready
	2. 如果多个channel同时ready，则随机选择一个执行
	3. 可以用于判断管道是否存满
*/

func test1(ch chan<- string){
	time.Sleep(5*time.Second)
	ch <- "test1"
}
func test2(ch chan<- string){
	time.Sleep(2*time.Second)
	ch <- "test2"
}

func test3(ch chan<- string)  {
	ch <- "test3"
}
func test4(ch chan<- string)  {
	ch <- "test4"
}

func write(ch chan string)  {
	for  {
		select {
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("chan full")
		}
		time.Sleep(time.Millisecond*500)
	}
}

func main() {
	//ch1 := make(chan string)
	//ch2 := make(chan string)

	//select可以同时监听一个或多个channel，直到其中一个channel ready
	//go test1(ch1)
	//go test2(ch2)
	//select {
	//case x:= <-ch1:
	//	fmt.Println("test1 result:", x)
	//case x := <-ch2:
	//	fmt.Println("test2 result:", x)
	//}

	// 如果多个channel同时ready，则随机选择一个执行
	//go test3(ch1)
	//go test4(ch2)
	//select {
	//case x:= <- ch1:
	//	fmt.Println(x)
	//case x:= <-ch2:
	//	fmt.Println(x)
	//}

	// 可以用于判断管道是否存满
	write1 := make(chan string)
	go write(write1)
	for str := range write1 {
		fmt.Println(str)
		time.Sleep(time.Second)
	}
}
