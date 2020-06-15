package main

import "fmt"

// Go语言中推荐使用驼峰式命名

var name string
var age int
var isOk bool
var array [5]int
var slice []int
var m map[string]int
var ch chan int

// 批量申明
var (
	n string
	i int
)

var s1 string = "whb"
var s2 = "20"

func main() {
	// 声明变量同时赋值
	var s1 string = "whb"
	fmt.Println(s1)
	// 类型推导（根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	// 简短变量声明，只能在函数里面用
	s3 := "哈哈哈"
	fmt.Println(s3)
}
