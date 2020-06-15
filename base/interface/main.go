package main

import "fmt"

// 接口
//接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
//在Go语言中接口（interface）是一种类型，一种抽象的类型。
//
//interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则），不关心属性（数据），只关心行为（方法）。
//接口是一个或多个方法签名的集合。
//    任何类型的方法集中只要拥有该接口'对应的全部方法'签名。
//    就表示它 "实现" 了该接口，无须在该类型上显式声明实现了哪个接口。
//    这称为Structural Typing。
//    所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值。
//    当然，该类型还可以有其他方法。
//
//    接口只有方法声明，没有实现，没有数据字段。
//    接口可以匿名嵌入其他接口，或嵌入到结构中。
//    对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。
//    只有当接口存储的类型和对象都为nil时，接口才等于nil。
//    接口调用不会做receiver的自动转换。
//    接口同样支持匿名字段方法。
//    接口也可实现类似OOP中的多态。
//    空接口可以作为任何类型数据的容器。
//    一个类型可实现多个接口。
//    接口命名习惯以 er 结尾。

// 1. 接口的定义
/*
type 接口类型名 interface{
	方法名1( 参数列表1 ) 返回值列表1
	方法名2( 参数列表2 ) 返回值列表2
	…
}
1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
2.方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
3.参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。
*/

// 2. 接口实现
type sayer interface {
	say()
}

type dog struct {

}
type cat struct {

}
func (d dog) say(){
	fmt.Println("汪汪汪～")
}
func (c cat) say(){
	fmt.Println("喵喵喵～")
}
func da(s sayer) {
	s.say()
}

// 3. 值接收者和指针接收者实现接口的区别
type Mover interface {
	move()
}
type person struct {

}
// 值接收者实现接口
/*func(p person) move(){
	fmt.Println("running")
}*/

//指针接收者实现接口
func (p *person) move()  {
	fmt.Println("running")
}

// 4. 类型与接口的关系
// 4.1 一个类型实现多个接口
// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现
// 见文件 interface_struct.go

// 6. 空接口
// 6.1 空接口的定义
//空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
//空接口类型的变量可以存储任意类型的变量。
var x interface{}

// 6.2 空接口的应用
//空接口作为函数的参数
//使用空接口实现可以接收任意类型的函数参数。
func show(a interface{}){
	fmt.Printf("type: %T  value: %v \n", a, a)
}
//空接口作为map的值
//使用空接口实现可以保存任意值的字典
var info map[string]interface{}

// 7. 类型断言
//接口值
//一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。
//想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：
//    x.(T)
//其中：
//    x：表示类型为interface{}的变量
//    T：表示断言x可能是的类型。
//该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
func justifyType(x interface{}){
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, val: %v \n", v)
	case int:
		fmt.Printf("x is a int , val: %v\n", v)
	case bool:
		fmt.Printf("x is a bool , val: %v \n",v)
	default:
		fmt.Printf("unsupprot type!")
	}
}
func main() {
	var d dog
	var c cat
	da(d)
	da(c)
	var s sayer   // 接口类型变量能够存储所有实现了该接口的实例
	s = d
	s = c
	fmt.Println(s)

	//值接收者和指针接收者实现接口的区别
	var x Mover			// 接口类型变量
	//var p1 = person{}
	//var p2 = &person{}
	////值接收者
	//x = p1				//x可以接收person类型
	//x = p2				//x可以接收*person类型
	//x.move()			//Go语言中有对指针类型变量求值的语法糖

	//指针接收者
	var p3 = &person{}
	x = p3
	x.move()

	// 类型断言
	var aaa interface{}
	aaa = "test"
	v, ok := aaa.(string)
	if !ok {
		fmt.Println("aaa is not string")
	}else {
		fmt.Println(v)
	}
}
