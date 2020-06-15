package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeByOS()  {
	fileObj, err := os.OpenFile("./xxx.data", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed , error:%v\n", err)
		return
	}
	fileObj.Write([]byte("test test test\n"))
	fileObj.WriteString("dgfsg demo")
	fileObj.Close()
}

func writeByBufio()  {

	fileObj, err := os.OpenFile("./xxx.data", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed , error:%v\n", err)
		return
	}
	defer fileObj.Close()
	write := bufio.NewWriter(fileObj)
	write.WriteString("demo demo \n")
	write.Write([]byte("test tset"))
	write.Flush()
}

func writeByIoutil()  {
	str := "hello \n"
	err := ioutil.WriteFile("./xxxx.data", []byte(str), 0666)
	if err != nil {
		 fmt.Printf("write file failed, error:%v", err)
		return
	}
}

func main() {
	
}
