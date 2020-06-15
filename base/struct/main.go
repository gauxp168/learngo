package main

import (
	"encoding/json"
	"fmt"
)

// type 类型

// 类型定义
type newInt int
// 类型别名
type myInt = int

// struct 结构体
//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）
// 1. 结构体定义
/*
type 类型名 struct {
	字段名 字段类型
	字段名 字段类型
	…
}
1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
3.字段类型：表示结构体字段的具体类型。
*/
type person struct {
	name string
	age int8
	city string
}

type x struct {
	a int8
	b int8
	c int8
}

// 构造函数
// 构造函数:约定成俗用new开头
// 返回的是结构体还是结构体指针
// 当结构体比较大的时候尽量使用结构体指针,减少程序的内存开销
func newPerson(name,city string, age int8) *person{
	return &person{
		name: name,
		age:age,
		city:city,
	}
}


// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量,多用类型名首字母小写表示
// 使用值接收者:传拷贝进去
/*
方法和接收者

值类型的接收者
当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	函数体
}
1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
3.方法名、参数列表、返回参数：具体格式与函数定义相同。

指针类型的接收者
// SetAge 设置p的年龄
// 使用指针接收者
指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}
什么时候应该使用指针类型接收者
1.需要修改接收者中的值
2.接收者是拷贝代价比较大的大对象
3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/
func (p person) newYear()  {
	p.age++
}
// 指针接收者:传内存地址进去
func (p *person) newYearP()  {
	p.age++
}

// 结构体嵌套
type address struct {
	city string
	province string
}
type company struct {
	name string
	address
}

// 结构体方法继承
type animal struct {
	name string
}

type dog struct {
	feet int8
	animal
}

func (a animal) move()  {
	fmt.Printf("%s会动！\n", a.name)
}

func (d dog) wang()  {
	fmt.Printf("%s再叫：王阿旺\n", d.name)//d.animal.name
}

// 给自定义类型加方法
// 不能给别的包里面的类型添加方法,只能给自己包里的类型添加方法
func (n newInt) add(){
	fmt.Println("this is int")
}

// // 结构体与json
//
//// 1.序列化:   把Go语言中的结构体变量 --> json格式的字符串
//// 2.反序列化: json格式的字符串   --> Go语言中能够识别的结构体变量
type student struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age int8 `json:"age"`
}

func main() {
	// 类型定义和类型别名的区别
	var a newInt
	var b myInt
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	// 结构体实例化
	//只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
	//结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
	// var 结构体实例 结构体类型
	var p person
	p.name = "tset"
	p.age = 23
	p.city = "shenzhen"

	// 匿名结构体
	var user struct{name string; age int8}
	user.name = "username"
	user.age = 22

	// 结构体占用一块连续的内存空间
	m := x{
		a: 10,
		b: 20,
		c: 30,
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))

	// 结构体 继承调用
	d1 := dog{
		feet:4,
		animal:animal{
			name:test,
		},
	}
	d1.move()
	d1.wang()

	//序列化
	stu := student{
		Name:"hello",
		Age:22,
	}
	jb,err := json.Marshal(stu)
	if err != nil {
		fmt.Printf("marshal failed err:%v\n", err)
		return
	}
	fmt.Printf("%v\n", string(jb))
	//反序列化
	str := `{"name":"理想","age":18}`
	var s2 student
	json.Unmarshal([]byte(str), &s2)
	fmt.Printf("%v\n", s2)

}
