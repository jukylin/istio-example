package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello V1")
}

func main() {
	http.HandleFunc("/", sayhello)       //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}