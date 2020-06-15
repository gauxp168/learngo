package main

// sync
// 1. sync.WaitGroup
/*
(wg * WaitGroup) Add(delta int)	计数器+delta
(wg *WaitGroup) Done()	计数器-1
(wg *WaitGroup) Wait()	阻塞直到计数器变为0
*/

// 2. sync.Once
/*
在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。
Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。
sync.Once只有一个Do方法，其签名如下：
func (o *Once) Do(f func()) {}
注意：如果要执行的函数f需要传递参数就需要搭配闭包来使用。
*/
//var icons map[string]image.Image
//var loadIconsOnce sync.Once
//
//func loadIcons()  {
//	icons = map[string]image.Image{
//		"left": loadIcon("left.png"),
//		"up": loadIcon("left.png"),
//		"right": loadIcon("left.png"),
//		"down": loadIcon("left.png"),
//	}
//}
//func Icon(name string) image.Image {
//	loadIconsOnce.Do(loadIcons)
//	return icons[name]
//}

// 3. sync.Map
// Go语言中内置的map不是并发安全的
// Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。
// 开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。
// 同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法

func main() {
	
}
