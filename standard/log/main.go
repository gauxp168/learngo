package main

import (
	"fmt"
	"log"
	"os"
)

// logger
// log包定义了Logger类型，该类型提供了一些格式化输出的方法。
// 本包也提供了一个预定义的“标准”logger，可以通过调用函数Print系列(Print|Printf|Println）
// 、Fatal系列（Fatal|Fatalf|Fatalln）、
// 和Panic系列（Panic|Panicf|Panicln）来使用，比自行创建一个logger对象更容易使用。

func logDemo()  {
	log.Println("test println")
	log.Panicln("test panic")
	log.Fatalln("test fata")
}

// 配置logger
//默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，
// 比如记录该日志的文件名和行号等。log标准库中为我们提供了定制这些设置的方法。
//
//log标准库中的Flags函数会返回标准logger的输出配置，而SetFlags函数用来设置标准logger的输出配置。
//    func Flags() int
//    func SetFlags(flag int)

// flag选项
//log标准库提供了如下的flag选项，它们是一系列定义好的常量。
//
/*const (
   // 控制输出日志信息的细节，不能控制输出的顺序和格式。
   // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
   Ldate         = 1 << iota     // 日期：2009/01/23
   Ltime                         // 时间：01:23:23
   Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
   Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
   Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
   LUTC                          // 使用UTC时间
   LstdFlags     = Ldate | Ltime // 标准logger的初始值
)*/

func flagDemo()  {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}

// 配置日志前缀
//log标准库中还提供了关于日志信息前缀的两个方法：
//
//    func Prefix() string
//    func SetPrefix(prefix string)
//其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀

func perfixDemo()  {
	log.SetFlags(log.Llongfile | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[pprof]")
	log.Println("这是一条很普通的日志。")
}

//  配置日志输出位置
//    func SetOutput(w io.Writer)
//SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出

func init()  {
	fileObj, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed, error:", err)
		return
	}
	log.SetOutput(fileObj)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

// 创建logger
//log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。New函数的签名如下：
//
//    func New(out io.Writer, prefix string, flag int) *Logger
//New创建一个Logger对象。其中，参数out设置日志信息写入的目的地。
// 参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）

func newDemo()  {
	logger := log.New(os.Stdout, "{new}", log.Llongfile|log.Lmicroseconds|log.Ldate)
	logger.Println("new logger test ")
}

//  Go内置的log库功能有限，例如无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap

func main() {
	//logDemo()
	// perfixDemo()
	newDemo()
}
