package main

// context

// 为什么需要Context?如何优雅的实现结束子goroutine?如何接收外部命令实现退出?
// a. 全局变量方式
// 全局变量方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易统一
// 2. 如果worker中再启动goroutine，就不太好控制了。

// b. 通道方式
// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel

// c. 官方版的方案
// 标准库context，它定义了Context类型，
// 专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

// 对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。
// 它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。
// 当一个上下文被取消时，它派生的所有上下文也被取消。

/*
Context接口
context.Context是一个接口，该接口定义了四个需要实现的方法。具体签名如下：

type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
其中：

Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）；
Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；
Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
如果当前Context被取消就会返回Canceled错误；
如果当前Context超时就会返回DeadlineExceeded错误；
Value方法会从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；
*/

/*
Background()和TODO()
Go内置两个函数：Background()和TODO()，这两个函数分别返回一个实现了Context接口的background和todo。我们代码中最开始都是以这两个内置的上下文对象作为最顶层的partent context，衍生出更多的子上下文对象。

Background()主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。

TODO()，它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。

background和todo本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。
*/

/*
With系列函数
此外，context包中还定义了四个With系列函数。

1. WithCancel
WithCancel的函数签名如下：
    func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
WithCancel返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况。
取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

2. WithDeadline
WithDeadline的函数签名如下：
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
返回父上下文的副本，并将deadline调整为不迟于d。如果父上下文的deadline已经早于d，则WithDeadline(parent, d)在语义上等同于父上下文。当截止日过期时，当调用返回的cancel函数时，或者当父上下文的Done通道关闭时，返回上下文的Done通道将被关闭，以最先发生的情况为准。
取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

3. WithTimeout
WithTimeout的函数签名如下：
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
WithTimeout返回WithDeadline(parent, time.Now().Add(timeout))。
取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制

4. WithValue
WithValue函数能够将请求作用域的数据与 Context 对象建立关系。声明如下：
    func WithValue(parent Context, key, val interface{}) Context
WithValue返回父节点的副本，其中与key关联的值为val。
仅对API和进程间传递请求域的数据使用上下文值，而不是使用它来传递可选参数给函数。
所提供的键必须是可比较的，并且不应该是string类型或任何其他内置类型，以避免使用上下文在包之间发生冲突。
WithValue的用户应该为键定义自己的类型。
为了避免在分配给interface{}时进行分配，上下文键通常具有具体类型struct{}。
或者，导出的上下文关键变量的静态类型应该是指针或接口。

*/

/*
使用Context的注意事项
推荐以参数的方式显示传递Context
以Context作为参数的函数方法，应该把Context作为第一个参数。
给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO()
Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
Context是线程安全的，可以放心的在多个goroutine中传递
*/

func main() {
	
}
