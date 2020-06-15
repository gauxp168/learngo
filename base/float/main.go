package main

import "fmt"

func main() {
	//math.MaxFloat32
	f1 := 1.234
	// 默认Go语言中的小数都是float64类型
	fmt.Printf("%T\n", f1)

	f2 := float32(1.234)
	fmt.Printf("%T\n", f2)
	// float32类型的值不能直接复赋值给float64类型的变量
	//f1 != f2
}
