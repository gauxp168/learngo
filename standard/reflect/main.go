package main

import (
	"fmt"
	"reflect"
)

// reflect 反射
// 反射是指在程序运行期对程序本身进行访问和修改的能力

// 变量的内在机制
//变量包含类型信息和值信息 var arr [10]int arr[0] = 10
//类型信息：是静态的元信息，是预先定义好的
//值信息：是程序运行过程中动态改变的

// 反射的使用
//reflect包封装了反射相关的方法
//获取类型信息：reflect.TypeOf，是静态的
//获取值信息：reflect.ValueOf，是动态的

// 1. 空接口与反射
//反射可以在运行时动态获取程序的各种详细信息
// a. 反射获取interface类型信息
func reflect_type(a interface{})  {
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)
	k := t.Kind() // kind() 可以获取具体类型
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a is float64")
	case reflect.String:
		fmt.Println("a is string")
	}
}

// b. 反射获取interface值信息
func reflect_value(a interface{})  {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a is :", v.Float())
	}
}

// c. 反射修改值信息
func reflect_set_value(a interface{})  {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		// 反射修改值
		v.SetFloat(6.4)
		fmt.Println("a is :", v.Float())
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("a is case:", v.Elem().Float())
		fmt.Println(v.Pointer())
	}
}

// 2. 结构体与反射
// a. 查看类型、字段和方法
type User struct {
	Id int
	Name string
	Age int
}
func (u User) Hello()  {
	fmt.Println("Hello")
}
func Poni(o interface{})  {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()
	fmt.Println(t.NumField(), v.NumField())
	for i:=0; i< t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v \n", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		fmt.Println(t.Field(i), v.Field(i))
		val := v.Field(i).Interface()
		fmt.Println("val: ", val)
	}
	fmt.Println("===========方法=========")
	for i:=0; i< t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}
}

// b. 查看匿名字段
type Boy struct {
	User
	Addr string
}
func Anonymous()  {
	m := Boy{User{1, "zs", 20}, "bj"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	// Anonymous：匿名
	fmt.Printf("%#v\n", t.Field(0))
	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}

// c. 修改结构体的值
func SetValue(o interface{})  {
	v := reflect.ValueOf(o)
	// 获取指针指向的元素
	v=v.Elem()
	// 取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("kuteng")
	}
}

// d. 调用方法
func callMethod()  {
	u := User{1,"test test", 22}
	v := reflect.ValueOf(u)
	m := v.MethodByName("Hello")
	// 没参数的情况下：
	var args []reflect.Value
	// 有参数的情况下：构建一些参数
	//args := []reflect.Value{reflect.ValueOf("demo")}
	// 调用方法，需要传入方法的参数
	m.Call(args)
}

// f. 获取字段的tag
type student struct {
	Name string `json:"name" db:"name"`
}
func getTag()  {
	var s student
	v := reflect.ValueOf(&s)
	t := v.Type()
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("db"))
}

func main() {
	/*
	空接口与反射
	var x float64 = 3.4
	reflect_type(x)
	reflect_value(x)
	// 反射认为下面是指针类型，不是float类型
	//reflect_set_value(x) error
	reflect_set_value(&x)
	fmt.Println("main:", x)
	*/

	/*
	结构体与反射
	u := User{1, "zs", 20}
	Poni(u)
	Anonymous()
	SetValue(&u)
	fmt.Println(u)
	getTag()*/
}
