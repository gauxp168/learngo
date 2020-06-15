package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Template
// html/template包实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出。
// 它提供了和text/template包相同的接口，Go语言中输出HTML的场景都应使用text/template包。

func sayHello(w http.ResponseWriter, r *http.Request)  {
	tmpl, err := template.ParseFiles("hello.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, "tests")
}

// 模板语法
//{{.}}
//模板语法都包含在{{和}}中间，其中{{.}}中的点表示当前对象。
//当我们传入一个结构体对象时，我们可以根据.来访问结构体的对应字段
// 同理，当我们传入的变量是map时，也可以在模板文件中通过.根据key来取值。
type User struct {
	Name string
	Age int
	Gender string
}
func sayhello1(w http.ResponseWriter, r *http.Request)  {
	tmpl, err := template.ParseFiles("./hello1.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	user := User{
		Name:"test",
		Age:22,
		Gender:"man",
	}
	tmpl.Execute(w, user)
}

// 注释
//    {{/* a comment */}}
//    注释，执行时会忽略。可以多行。注释不能嵌套，并且必须紧贴分界符始止。


// pipeline
//pipeline是指产生数据的操作。比如{{.}}、{{.Name}}等。
// Go的模板语法中支持使用管道符号|链接多个命令，
// 用法和unix下的管道类似：|前面的命令会将运算结果(或返回值)传递给后一个命令的最后一个位置。
//
//注意 : 并不是只有使用了|才是pipeline。
// Go的模板语法中，pipeline的概念是传递数据，只要能产生数据的，都是pipeline。

// 变量
//Action里可以初始化一个变量来捕获管道的执行结果。初始化语法如下：
//
//    $variable := pipeline
//其中$variable是变量的名字。声明变量的action不会产生任何输出

// 条件判断
//Go模板语法中的条件判断有以下几种:
//{{if pipeline}} T1 {{end}}
//{{if pipeline}} T1 {{else}} T0 {{end}}
//{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}


// range
//Go的模板语法中使用range关键字进行遍历，有以下两种写法，
// 其中pipeline的值必须是数组、切片、字典或者通道。

//{{range pipeline}} T1 {{end}}
//如果pipeline的值其长度为0，不会有任何输出
//{{range pipeline}} T1 {{else}} T0 {{end}}
//如果pipeline的值其长度为0，则会执行T0。

// with
//{{with pipeline}} T1 {{end}}
//如果pipeline为empty不产生输出，否则将dot设为pipeline的值并执行T1。不修改外面的dot。
//
//{{with pipeline}} T1 {{else}} T0 {{end}}
//如果pipeline为empty，不改变dot并执行T0，否则dot设为pipeline的值并执行T1。

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/hello", sayhello1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err:", err)
		return
	}
}
