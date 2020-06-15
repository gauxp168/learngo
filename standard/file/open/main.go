package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFileByOS(){
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed error:%v\n",err)
		return
	}
	// 关闭文件
	defer fileObj.Close()
	// 读取文件
	var tmp [128]byte
	for  {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}

		if err != nil {
			fmt.Printf("read file failed, error:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:]))
		if n < 128 {
			return
		}
	}
}

func readFromFilebyBufio(){
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("file open failed, error:%v\n", err)
		return
	}
	defer  fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for  {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, error:%v\n", err)
			return
		}
		fmt.Println(line)
	}
}

func readFromFileByIoutil()  {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed, error:%v\n",err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	// 1. 文件对象的类型
	// 2. 获取文件对象的详细信息
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, error:%v\n", err)
	}
	fmt.Printf("fileObj type %T\n", fileObj)
	infoObj, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file status info failed, error :%V\n", err)
	}
	infoObj.Size()
	infoObj.Name()
}
