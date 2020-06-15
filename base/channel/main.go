package main

import (
	"fmt"
	"sync"
)

// channel
//Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
//channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。
//Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

//channel是一种类型，一种引用类型。声明通道类型
//var 变量 chan 元素类型

// 1. 定义 channel
var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递int切片的通道
var wg sync.WaitGroup

// 4. 无缓冲的通道
// 无缓冲的通道又称为阻塞的通道
// 使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道

func noBufChannel(){
	ch1 = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch1
		fmt.Println("goroutine 从通道ch1中取到数据",x)
	}()
	ch1 <- 10
	wg.Wait()
}

// 5. 有缓冲的通道
// 只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量
// 我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量
func bufChannel(){
	ch1 = make(chan int , 10)
	ch1<- 10
	ch1<- 20
	x := <- ch1
	fmt.Println(x)
}

//  6. 单向通道
// 有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。
// 其中，
//
//    1.chan<- int是一个只能发送的通道，可以发送但是不能接收；
//    2.<-chan int是一个只能接收的通道，可以接收但是不能发送。
func counter( out chan<- int)  {
	for i:= 0; i< 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int)  {
	for key := range in {
		out <- key*key
	}
	close(out)
}

func printer(in <-chan int)  {
	for key := range in {
		fmt.Println(key)
	}
}

func main() {
	// 通道是引用类型，通道类型的空值是nil。
	fmt.Println(ch1)
	// 2. 创建 channel
	// make(chan 元素类型, [缓冲大小])
	ch4 := make(chan int)
	ch5 := make(chan int, 10)
	// 3. 通道有发送（send）、接收(receive）和关闭（close）三种操作。
	// 发送:将一个值发送到通道中
	ch4<- 10
	ch5<- 10
	//接收：从一个通道中接收值
	x := <- ch4
	fmt.Println(x)
	<- ch5
	// 关闭： 我们通过调用内置的close函数来关闭通道
	close(ch4)
	close(ch5)
	//关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
	/*
	关闭后的通道有以下特点：
    1.对一个关闭的通道再发送值就会导致panic。
    2.对一个关闭的通道进行接收会一直获取值直到通道为空。
    3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
    4.关闭一个已经关闭的通道会导致panic。
	*/

	// 如何判断一个通道是否被关闭
	/*
	1.   i, ok := <-ch1 // 通道关闭后再取值ok=false
	2.   for i := range ch2 { // 通道关闭后会退出for range循环
        fmt.Println(i)
    }
	*/

}
