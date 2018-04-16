package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sleep := r.Form.Get("sleep")
	resp, err := http.Get("http://sleep:8080?sleep=" + sleep)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(body))

	fmt.Fprint(w, "Hello V1")
}

func main() {
	http.HandleFunc("/", sayhello)       //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}