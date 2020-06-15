package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/go", myHandler)
	http.ListenAndServe("127.0.0.1:8000", nil)
}

func myHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println(r.RemoteAddr, "连接成功")
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	fmt.Println(r.Header)
	fmt.Println(r.Body)

	w.Write([]byte("how are you?"))
}