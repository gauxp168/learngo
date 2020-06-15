package main

import (
	"fmt"
	"io"
	"os"
)

// 在文件中插入数据案例

// 1. defer close 的位置问题
// f11 函数位置是错误的
func f1() {
	var fileObj *os.File
	var err error
	fileObj,err = os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed , error:%v\n", err)
		return
	}
	defer fileObj.Close()
}
func f11() {
	var fileObj *os.File
	var err error
	fileObj,err = os.Open("./main.go")
	defer fileObj.Close()
	if err != nil {
		fmt.Printf("open file failed , error:%v\n", err)
		return
	}
}

func fileInsert() {
	fileObj, err := os.OpenFile("./sb.data", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, error:%v\n", err)
		return
	}
	defer fileObj.Close()
	// 因为没有办法直接在文件中间插入内容，所以要借助一个临时文件
	tmpObj,err := os.OpenFile("sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open tmp file failed, error:%v\n", err)
		return
	}
	defer tmpObj.Close()

	// 1. 读取源文件写到临时文件
	var first [2]byte
	n, err := fileObj.Read(first[:])
	if err != nil {
		fmt.Printf("read file failed, error:%v\n", err)
		return
	}
	tmpObj.Write(first[:n])
	// 插入新内容
	var c  = []byte{'C'}
	tmpObj.Write(c)
	// 最后把源文件后续的内容写入临时文件
	var s  [1024]byte
	for  {
		n, err := fileObj.Read(s[:])
		if err == io.EOF {
			tmpObj.Write(s[:n])
			break
		}
		if err != nil {
			fmt.Printf("read file failed, error:%v\n", err)
			return
		}
		tmpObj.Write(s[:n])
	}
	os.Rename("./sb.tmp", "./sb.data")

}

func main() {
	
}
