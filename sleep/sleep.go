package main

import (
	"log"
	"net/http"
	"time"
	"strconv"
	"net"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sleep := r.Form.Get("sleep")
	sec,_ := strconv.Atoi(sleep)
	time.Sleep(time.Duration(sec) * time.Second)
	w.Write([]byte("Sleep "+ sleep + " ip " + getPrivateIPIfAvailable().String() + "\n"))
}

func retry(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	status := r.Form.Get("status")
	s ,_ := strconv.Atoi(status)

	w.WriteHeader(s)
	w.Write([]byte( " rep status" + status + " ip " + getPrivateIPIfAvailable().String() + "\n"))
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
	mux := http.NewServeMux()

	mux.HandleFunc("/", sayhello)
	mux.HandleFunc("/retry", retry)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}