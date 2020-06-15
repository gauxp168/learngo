package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// IO

// 输入输出的底层原理
//终端其实是一个文件，相关实例如下：
//os.Stdin：标准输入的文件实例，类型为*File
//os.Stdout：标准输出的文件实例，类型为*File
//os.Stderr：标准错误输出的文件实例，类型为*File

// 文件操作相关API
//func Create(name string) (file *File, err Error)
//根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666
//func NewFile(fd uintptr, name string) *File
//根据文件描述符创建相应的文件，返回一个文件对象
//func Open(name string) (file *File, err Error)
//只读方式打开一个名称为name的文件
//func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
//打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限
//func (file *File) Write(b []byte) (n int, err Error)
//写入byte类型的信息到文件
//func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
//在指定位置开始写入byte类型的信息
//func (file *File) WriteString(s string) (ret int, err Error)
//写入string信息到文件
//func (file *File) Read(b []byte) (n int, err Error)
//读取数据到b中
//func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
//从off开始读取数据到b中
//func Remove(name string) Error
//删除文件名为name的文件

// 打开和关闭文件
//os.Open()函数能够打开一个文件，返回一个*File和一个err。
// 对得到的文件实例调用close()方法能够关闭文件。
func openFile1()  {
	file, err := os.Open("./adv.data")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
}

// 写文件

func writeFile()  {
	file, err := os.Create("./adv.data")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.Write([]byte("abc"))
	file.WriteString("def")
}

// 读文件
//文件读取可以用file.Read()和file.ReadAt()，读到文件末尾会返回io.EOF的错误
func readFile()  {
	file, err := os.Open("./adv.data")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var buf [128]byte
	var content []byte
	for  {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

// 拷贝文件
func copyFile()  {
	srcFile, err := os.Open("./adv.data")
	if err != nil {
		fmt.Println(err)
		return
	}
	dstFile, err := os.Create("./new.data")
	if err != nil {
		fmt.Println(err)
		return
	}
	var buf [1024]byte
	for  {
		n, err := srcFile.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = dstFile.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
		}
	}
	srcFile.Close()
	dstFile.Close()
}

// bufio
//bufio包实现了带缓冲区的读写，是对文件读写的封装
//bufio缓冲写数据
/*
模式				含义
os.O_WRONLY		只写
os.O_CREATE		创建文件
os.O_RDONLY		只读
os.O_RDWR		读写
os.O_TRUNC		清空
os.O_APPEND		追加
*/

func bufiowr()  {
	// 参数2：打开模式，所有模式d都在上面
	// 参数3是权限控制
	// w写 r读 x执行   w  2   r  4   x  1
	file, err := os.OpenFile("./adv.data", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write([]byte("test"))
	writer.WriteString("demo")
	writer.Flush()
}

func bufiord()  {
	file, err := os.Open("adv.data")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//var buf [128]byte
	for  {
		//n, err := reader.Read(buf[:])
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(line))
	}
}

// ioutil工具包
//工具包写文件
//工具包读取文件

func ioutilwr()  {
	err := ioutil.WriteFile("adv.data", []byte("test tests"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ioutilre()  {
	bytes, err := ioutil.ReadFile("adv.data")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}

// 实现一个cat命令
//使用文件操作相关知识，模拟实现linux平台cat命令的功能
func main() {
	//var buf [16]byte
	//n, _ := os.Stdin.Read(buf[:])
	//os.Stdin.WriteString(string(buf[:n]))
	//fmt.Println(string(buf[:n]))

	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i:=0;i<flag.NArg() ;i++  {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		defer file.Close()
		cat(bufio.NewReader(file))
	}
}

func cat(r *bufio.Reader)  {
	for  {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stdout, "%s", err)
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}
