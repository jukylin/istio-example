package main

import (
	"log"
	"net/http"
	"net"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := r.Form.Get("panic")
	if p == "1"{
		panic("pance 退出")
	}


	w.Write([]byte("Hello V2 \n"))
	w.Write([]byte("ip " + getPrivateIPIfAvailable().String() + "\n"))
}

func getPrivateIPIfAvailable() net.IP {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		default:
			continue
		}
		if !ip.IsLoopback() {
			return ip
		}
	}
	return net.IPv4zero
}


func main() {
	http.HandleFunc("/", sayhello)       //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}