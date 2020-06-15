package main

import "fmt"

// 4. 类型与接口的关系
// 4.1 一个类型实现多个接口
// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现

// Sayer 接口
type Sayer interface {
	say()
}

// Runer 接口
type Runer interface {
	run()
}

type god struct {
	name string
}

type car struct {
	brand string
}

func (g god) say(){
	fmt.Println("主呀")
}

func (g god) run(){
	fmt.Println("飞")
}

// 4.2 多个类型实现同一接口
// Go语言中不同的类型还可以实现同一接口

func (c car) run(){
	fmt.Println("车")
}
//并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
type WashingMachine interface {
	wash()
	dry()
}
type dryer struct{

}
func (d dryer) dry(){
	fmt.Println("睡衣睡")
}
type haier struct {
	dryer
}
func (h haier) wash(){
	fmt.Println("洗刷刷")
}


// 5. 接口嵌套
// 接口与接口间可以通过嵌套创造出新的接口
type midea struct{
	dryer
}

func main() {
	
}
