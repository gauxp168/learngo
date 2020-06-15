package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		fmt.Println("http get failed, error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)
	for  {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}else{
			fmt.Println("读取完毕")
			fmt.Println(string(buf[:n]))
			break
		}
	}
}
