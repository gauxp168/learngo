package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	err := http.ListenAndServe("127.0.0.1:8080", router)
	if err != nil {
		fmt.Println("error:",err)
	}
}
