package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"strconv"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sleep := r.Form.Get("sleep")
	sec,_ := strconv.Atoi(sleep)
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Fprint(w, "Sleep "+ sleep)
}

func main() {
	http.HandleFunc("/", sayhello)       //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}