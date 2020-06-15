package main

import "fmt"

// 指针
func main() {
	// 1. & 取地址符
	n := 8
	p := &n
	fmt.Println(p)
	fmt.Printf(" %T\n",  p)

	// 2. * 根据地址取值
	m:= *p
	fmt.Println(m)
	fmt.Printf("%T \n", m)

	// 定义指针类型变量
	var a1 *int
	fmt.Println(a1)
	var a2 = new(int)  // new 函数申请一个内存地址
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)

	a := 100
	b := &a
	fmt.Printf("type a:%T type b:%T\n", a,b)
	fmt.Printf("&a %p\n", &a)
	fmt.Printf("bp %p\n", b)
	fmt.Printf("bv %v\n", b)
	fmt.Printf("&b %p\n", &b)
}
