package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// flag
// Go语言内置的flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单

// os.Args
// 如果你只是简单的想要获取命令行参数，可以像下面的代码示例一样使用os.Args来获取命令行参数
// os.Args是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称
func osArgs()  {
	if len(os.Args) > 0 {
		for k, v := range os.Args {
			fmt.Printf("args[%d]=%v\n", k, v)
		}
	}
}

// 1. flag参数类型
//flag包支持的命令行参数类型有bool、int、int64、uint、uint64、float float64、string、duration。
//
//flag参数		有效值
//字符串flag		合法字符串
//整数flag		1234、0664、0x1234等类型，也可以是负数。
//浮点数flag		合法浮点数
//bool类型flag		1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False。
//时间段flag		任何合法的时间段字符串。如”300ms”、”-1.5h”、”2h45m”。
//合法的单位有”ns”、”us” /“µs”、”ms”、”s”、”m”、”h”。


func main() {
	// 2. 定义命令行flag参数
	//有以下两种常用的定义命令行flag参数的方法。
	//
	// 2.1 flag.Type()
	// 基本格式如下：
	//flag.Type(flag名, 默认值, 帮助信息)*Type
	name1 := flag.String("name", "test", "姓名")
	age1 := flag.Int("age", 32, "年龄")
	married1 := flag.Bool("married", true, "婚否")
	delay1 := flag.Duration("delay", 1, "时间间隔")
	fmt.Printf("name:%s age:%d married:%t delay：%v\n", name1, age1, married1, delay1)
	// 需要注意的是，此时name、age、married、delay均为对应类型的指针

	// 2.2 flag.TypeVar()
	//基本格式如下： flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
	var (
		name string
		age int
		married bool
		delay time.Duration
	)
	flag.StringVar(&name, "name", "test", "姓名")
	flag.IntVar(&age, "age", 33, "年龄")
	flag.BoolVar(&married, "married", true, "婚否")
	flag.DurationVar(&delay, "defay", 1, "时间间隔")

	// 2.3 flag.Parse()
	//通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。
	//
	//支持的命令行参数格式有以下几种：
	//
	//-flag xxx （使用空格，一个-符号）
	//--flag xxx （使用空格，两个-符号）
	//-flag=xxx （使用等号，一个-符号）
	//--flag=xxx （使用等号，两个-符号）
	//其中，布尔类型的参数必须使用等号的方式指定。
	//
	//Flag解析在第一个非flag参数（单个”-“不是flag参数）之前停止，或者在终止符”–“之后停止。

	// 2.4 flag其他函数
	//flag.Args() ////返回命令行参数后的其他参数，以[]string类型
	//flag.NArg() //返回命令行参数后的其他参数个数
	//flag.NFlag() //返回使用的命令行参数个数

	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}
/*
使用
命令行参数使用提示：

    $ ./flag_demo -help
    Usage of ./flag_demo:
      -age int
            年龄 (default 18)
      -d duration
            时间间隔
      -married
            婚否
      -name string
            姓名 (default "张三")
正常使用命令行flag参数：

    $ ./flag_demo -name pprof --age 28 -married=false -d=1h30m
    pprof 28 false 1h30m0s
    []
    0
    4
使用非flag命令行参数：

    $ ./flag_demo a b c
    张三 18 false 0s
    [a b c]
    3
    0
*/