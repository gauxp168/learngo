package main

import (
	"fmt"
	"strings"
)

// 函数

// 函数存在的意义？
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中，给它起个名字，每次用到它的时候直接用函数名调用就可以了
// 使用函数能够让代码结构更清晰、更简洁。

// 函数定义
func sum(x int, y int) int {
	return x + y
}

// 返回值可以命名也可以不命名
// 命名的返回值就相当于在函数中声明一个变量
func sum11(x int, y int) (ret int) {
	ret =  x + y
	// 使用命名返回值可以return后省略
	return
}

// 多个返回值
func f()(int, string)  {
	return 1, "hello"
}

// 参数的类型简写:
// 当参数中连续多个参数的类型一致时，我们可以将非最后一个参数的类型省略
func f1(x,y,z int, m,n string) int{
	return  x+y+z
}

// 可变长参数
// 可变长参数必须放在函数参数的最后
func f2(x string, y ...int)  {
	fmt.Println(x)
	fmt.Println(y)// y的类型是切片 []int
}
//可变参数类型
func f3 (m string, n ...interface{}){
	fmt.Println(m)
	fmt.Println(n)
}

//defer
// defer多用于函数结束之前释放资源（文件句柄、数据库连接、socket连接）
func deferDemo(){
	fmt.Println("start")
	//defer把它后面的语句延迟到函数即将返回的时候再执行
	// 一个函数中可以有多个defer语句
	// 多个defer语言按照先进后出（后进先出）的顺序延迟执行
	defer fmt.Println("hello")
	defer fmt.Println("test")
	defer fmt.Println("demo")
}

// Go语言中函数的return不是原子操作，在底层是分为两步来执行
// 第一步：返回值赋值
// defer
// 第二步：真正的RET返回
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间
func df1() int{
	x:= 5
	defer func() {
		x++   // 修改的是x不是返回值
	}()
	return  x   // 1. 返回值赋值 2. defer 3. 真正的RET指令
}

func df2() (x int) {
	defer func() {
		x++    // 修改的是返回值
	}()
	return 5   // 1. 返回值赋值 2. defer 3. 真正的RET指令
}

func def (index string, a,b int) int{
	ret:= a+b
	fmt.Println(index, a,b,ret)
	return ret
}
/*
#调用
a := 1
b := 2
defer calc("1", a, calc("10", a, b))
a = 0
defer calc("2", a, calc("20", a, b))
#结果
1. a:=1
2. b:=2
3. defer calc("1", 1, calc("10", 1, 2))
4. calc("10", 1, 2) // "10" 1 2 3
5. defer calc("1", 1, 3)
6. a = 0
7. defer calc("2", 0, calc("20", 0, 2))
8. calc("20", 0, 2) // "20" 0 2 2
9. defer calc("2", 0, 2)
10. b = 1
calc("2", 0, 2) // "2" 0 2 2
calc("1", 1, 3) // "1" 1 3 4
*/

// 函数中查找变量的顺序
// 1. 先在函数内部查找
// 2. 找不到就往函数的外面查找,一直找到全局



// 函数类型
// 函数也可以作为参数的类型
func fp(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func ff(a, b int) int  {
	return a+b
}

func fp1(x func() int) func(int, int) int{
	return  ff
}

// 闭包
// 把原来需要传递两个int类型的参数包装成一个不需要传参的函数
func fb1(x,y int) int{
	return  x+y
}
func fb2(f func(int, int) int, x, y int)func(){
	return func() {
		f(x , y)
	}
}
// 闭包是什么？
// 闭包是一个函数，这个函数包含了他外部作用域的一个变量
// 底层的原理：
// 1. 函数可以作为返回值
// 2. 函数内部查找变量的顺序，先在自己内部找，找不到往外层找
func adder1() func(int) int{
	x := 100
	return func(y int) int {
		return x+y
	}
}

func adder2 (x int) func(int) int{
	return func(y int) int {
		return x+y
	}
}

func makeSuffix(suffix string) func(string)string{
	return func(str string) string {
		if !strings.HasSuffix(str, suffix) {
			return str + suffix
		}
		return str
	}
}

//关键字 函数名(参数)（返回值）{}
//函数调用影响基础参数
func calc(base int) (func(int) int, func(int) int){
	add := func(i int) int{
		base += i
		return base
	}

	sub := func(i int) int{
		base -= i
		return base
	}
	return add, sub
}
/*
注意：
子函数调用，影响基础参数数值
f1, f2 := calc(10)
fmt.Println(f1(1), f2(2)) // 11 9
fmt.Println(f1(3), f2(4)) // 12 8
fmt.Println(f1(5), f2(6)) // 13 7
*/


//// panic 和 recover
//Golang 没有结构化异常，使用 panic 抛出错误，recover 捕获错误。
//异常的使用场景简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
// 1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
//2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
//3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用
//延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
func funcB(){
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("test print")
	}()
	panic("this is error")
	fmt.Println("the end")
}

// 单行注释
/*
多行注释
*/
// Go语言函数外的语句必须以关键字开头
// main函数是入口函数
// 它没有参数也没有返回值
// 非main包可以没有main函数并且非main包里面的main函数非内部函数
func main() {
	// 函数内部定义的变量必须使用

	deferDemo()
}
